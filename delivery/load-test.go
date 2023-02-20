package delivery

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/hdkef/jameter/models"
)

func printLoadTestResult(req models.ReqsWrapper, counter int, timeTaken int64, totalReqs int, resultMap map[int]int) {
	fmt.Printf("URI\t\t\t: %s\n", req.URI)
	fmt.Printf("total requests\t\t: %d\n", counter)
	fmt.Printf("time taken\t\t: %d ms\n", timeTaken)
	fmt.Printf("throughput\t\t: %.3f req/s\n", 1000*float64(totalReqs)/float64(timeTaken))

	for k, v := range resultMap {
		fmt.Printf("Status Code %d total\t: %d\n", k, v)
	}

	fmt.Println("result may vary, make sure web servers don't limit req per ip")
}

func tagAsDone(resp *http.Response, wg *sync.WaitGroup, mtx *sync.Mutex, resultMap map[int]int, counter *int) {
	mtx.Lock()
	*counter++
	if *counter%100 == 0 {
		fmt.Println(*counter)
	}
	if resp.StatusCode != 0 {
		resultMap[resp.StatusCode]++
	}
	mtx.Unlock()
	wg.Done()
}

func byReqhit(r models.ReqsWrapper, wg *sync.WaitGroup, counter *int, mtx *sync.Mutex, client *http.Client, resultMap map[int]int) {
	//hit endpoint
	//create new request
	req, err := http.NewRequest(r.Method, r.URI, nil)

	if err != nil {
		// fmt.Println(err.Error())
		tagAsDone(&http.Response{}, wg, mtx, resultMap, counter)
		return
	}

	//add headers
	for _, v := range r.Headers {
		req.Header.Add(v.Name, v.Value)
	}

	//add cookies
	for _, v := range r.Cookies {
		cookies := http.Cookie{
			Name:  v.Name,
			Value: v.Value,
		}
		req.AddCookie(&cookies)
	}

	//add payload
	addPayload(req, &r)

	//execute the request
	resp, err := client.Do(req)
	if err != nil {
		// fmt.Println(err.Error())
		tagAsDone(&http.Response{}, wg, mtx, resultMap, counter)
		return
	}

	//done increment counter
	tagAsDone(resp, wg, mtx, resultMap, counter)
}

func loadTestByTotalReqs(project *models.Project, req models.ReqsWrapper) int {
	validTotalReq := false
	var totalReqs int

	//input total request
	for !validTotalReq {
		fmt.Print("\nInput total requests :")
		_, err := fmt.Scanln(&totalReqs)
		if err != nil {
			fmt.Println("Input invalid")
			continue
		}
		if totalReqs <= 0 {
			fmt.Println("Total req cannot be <= 0")
			continue
		}
		validTotalReq = true
	}

	startTime := time.Now()
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	client := &http.Client{}

	var counter int
	var resultMap map[int]int = make(map[int]int)

	//hit endpoints with goroutine
	for i := 0; i < totalReqs; i++ {
		wg.Add(1)
		go byReqhit(req, &wg, &counter, &mtx, client, resultMap)
	}

	//wait until all goroutine finished
	wg.Wait()

	timeTaken := time.Since(startTime).Milliseconds()

	printLoadTestResult(req, counter, timeTaken, totalReqs, resultMap)

	return 0
}

func byTimeResultListener(resultC chan *http.Response, doneC chan bool, mtx *sync.Mutex, counter *int, resultMap map[int]int) {
	isDone := false
out:
	for {
		select {
		case <-doneC:
			isDone = true
			break out
		case r := <-resultC:
			if isDone {
				//if is done is true, ignore any new response
				continue
			}
			mtx.Lock()
			*counter++
			// time.Sleep(1 * time.Millisecond)
			resultMap[r.StatusCode]++
			mtx.Unlock()
		}
	}
}

func byTimeHit(client *http.Client, r models.ReqsWrapper, resultC chan *http.Response) {
	//hit endpoint
	//create new request
	req, err := http.NewRequest(r.Method, r.URI, nil)

	if err != nil {
		// fmt.Println(err.Error())
		return
	}

	//add headers
	for _, v := range r.Headers {
		req.Header.Add(v.Name, v.Value)
	}

	//add cookies
	for _, v := range r.Cookies {
		cookies := http.Cookie{
			Name:  v.Name,
			Value: v.Value,
		}
		req.AddCookie(&cookies)
	}

	//add payload
	addPayload(req, &r)

	//execute the request
	resp, err := client.Do(req)
	if err != nil {
		// fmt.Println(err.Error())
		return
	}

	//send to chan
	resultC <- resp
}

func loadTestByTime(project *models.Project, req models.ReqsWrapper) int {
	validTotalTime := false
	totalTime := 0
	var resultC chan *http.Response = make(chan *http.Response)
	var doneC chan bool = make(chan bool)

	//input time in ms
	for !validTotalTime {
		fmt.Print("\nInput total time :")
		_, err := fmt.Scanln(&totalTime)
		if err != nil {
			fmt.Println("Input invalid")
			continue
		}
		validTotalTime = true
	}

	//validate total time
	if totalTime < 20000 {
		fmt.Println("Input minimal 20000 ms")
		return 0
	}

	//execute reqs
	client := &http.Client{}
	mtx := sync.Mutex{}
	var counter int
	var resultMap map[int]int = make(map[int]int)

	//create timer
	isDone := false
	time.AfterFunc(time.Duration(totalTime*int(time.Millisecond)), func() {
		fmt.Println("timer done")
		isDone = true
		doneC <- true
		//when timer finish, stop hitting endpoint and stop counting result
	})

	//listener
	go byTimeResultListener(resultC, doneC, &mtx, &counter, resultMap)

	for !isDone {
		//hit
		go byTimeHit(client, req, resultC)
	}

	//when done, tell listener to close chan
	printLoadTestResult(req, counter, int64(totalTime), counter, resultMap)
	// closeC <- true

	return 0
}

func LoadTest(project *models.Project) (menu int) {
	//shows all reqs
	ReadRequest(project)

	//Choose req ID
	var id int
	fmt.Print("\nInput request ID :")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Input invalid")
		return -1
	}

	//find req by ID
	var req models.ReqsWrapper
	for _, v := range project.Reqs {
		if v.ID == id {
			req = v
			break
		}
	}
	//check if req found
	if req.ID == 0 {
		fmt.Println("No request found")
		return -1
	}

	//Choose test type
	var opt int = -1
	for opt != 0 {
		fmt.Print("\n1.By total requests\n2.By time\n\nChoose menu :")
		_, err := fmt.Scanln(&opt)
		if err != nil {
			fmt.Println("Input invalid")
			continue
		}
		switch opt {
		case 1:
			opt = loadTestByTotalReqs(project, req)
		case 2:
			opt = loadTestByTime(project, req)
		}
	}
	return -1
}
