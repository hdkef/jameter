package delivery

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/hdkef/jameter/models"
)

func resChanListener(resultC chan interface{}, wg *sync.WaitGroup, closeC chan bool) {
	for {
		res := <-resultC
		fmt.Println("response :\n", res)
		wg.Done()
		if done := <-closeC; done {
			break
		}
	}
}

func addPayload(req *http.Request, r *models.ReqsWrapper) {

}

func hit(r *models.ReqsWrapper, resultC chan interface{}, wg *sync.WaitGroup, client *http.Client) {
	//create new request
	req, err := http.NewRequest(r.Method, r.URI, nil)

	if err != nil {
		wg.Done()
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
	addPayload(req, r)

	//execute the request
	resp, err := client.Do(req)
	if err != nil {
		wg.Done()
	}

	//decode response

	resultC <- resp
}

func ExecuteByIDS(project *models.Project) {

}

func ExecuteAll(project *models.Project) {
	var resultC chan interface{} = make(chan interface{})
	var doneC chan bool = make(chan bool)

	//create waitGroup
	wg := sync.WaitGroup{}

	//create http client
	client := &http.Client{}

	//execute every req
	for _, v := range project.Reqs {
		wg.Add(1)
		go hit(&v, resultC, &wg, client)
	}

	//listener
	go resChanListener(resultC, &wg, doneC)

	//wait for response completed
	wg.Wait()

	//if all executed, close listener
	doneC <- true
}

func ExecuteRequest(project *models.Project) (menu int) {
	fmt.Print("\n1.Execute all\n2.Execute by req ID\n\nChoose menu :")
	var opt int = -1
	_, err := fmt.Scanln(&opt)
	if err != nil {
		fmt.Println("Invalid input")
		return -1
	}

	switch opt {
	case 1:
		ExecuteAll(project)
	case 2:
		ExecuteByIDS(project)
	}

	return -1
}
