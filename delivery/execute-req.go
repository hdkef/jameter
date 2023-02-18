package delivery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hdkef/jameter/models"
	"github.com/hdkef/jameter/usecase"
)

func resChanListener(resultC chan *models.ResultChan, wg *sync.WaitGroup, closeC chan bool) {
out:
	for {
		select {
		case <-closeC:
			close(resultC)
			close(closeC)
			break out
		case res := <-resultC:

			var output string

			//handle response body according to MIME type
			defer res.Res.Body.Close()
			bodyBytes, err := io.ReadAll(res.Res.Body)
			if err != nil {
				fmt.Print(err.Error())
			}
			var body interface{}
			if res.Res.Header.Get("Accept") == "application/json" {
				var jsonBody map[string]interface{}
				err = json.Unmarshal(bodyBytes, &jsonBody)
				if err != nil {
					fmt.Print(err.Error())
				}
				body = jsonBody
			} else {
				body = string(bodyBytes)
			}

			output += fmt.Sprint("Date\t: ", time.Now().String(), "\n")
			output += fmt.Sprint("Status Code\t: ", res.Res.StatusCode, "\n")
			output += fmt.Sprint("Status Info\t: ", res.Res.Status, "\n")

			//output response header
			for k, v := range res.Res.Header {
				output += fmt.Sprintf("%s\t: %s\n", k, v)
			}

			output += fmt.Sprint("Body\t: \n", body, "\n")
			output += "\n\n\n"

			//append output to file

			fileName := fmt.Sprintf("%s_%s",
				res.Req.Name,
				res.Res.Request.URL.Host)

			fileName = strings.ReplaceAll(fileName, " ", "_")
			fileName += ".jamet"

			fileDir := fmt.Sprintf("output/%s", fileName)

			f, err := os.OpenFile(fileDir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Print(err.Error())
			}

			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				fmt.Print(err.Error())
			}

			//print result status
			fmt.Printf("\nStatus\t\t: %d\n", res.Res.StatusCode)
			fmt.Printf("Details on\t: %s\n", fileDir)

			wg.Done()
		}
	}
}

func addPayload(req *http.Request, r *models.ReqsWrapper) {

}

func hit(r models.ReqsWrapper, resultC chan *models.ResultChan, wg *sync.WaitGroup, client *http.Client) {
	//create new request
	req, err := http.NewRequest(r.Method, r.URI, nil)

	if err != nil {
		fmt.Println(err.Error())
		wg.Done()
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
		fmt.Println(err.Error())
		wg.Done()
		return
	}

	//decode response

	resultC <- &models.ResultChan{
		Res: resp,
		Req: &r,
	}

}

func ExecuteByIDS(project *models.Project) {
	var resultC chan *models.ResultChan = make(chan *models.ResultChan)
	var doneC chan bool = make(chan bool)

	//input ids to be executed
	prompter := usecase.Prompt{}
	idSlice, msg, v := prompter.GetReqIDSlice()
	if !v {
		fmt.Println(msg)
		return
	}

	//validate reqs
	if len(idSlice) == 0 {
		fmt.Println("No reqs found")
		close(resultC)
		close(doneC)
		return
	}

	//create waitGroup
	wg := sync.WaitGroup{}

	//create http client
	client := &http.Client{}

	//execute every req based on input
	for _, v := range idSlice {
		for _, k := range project.Reqs {
			if k.ID == v {
				wg.Add(1)
				go hit(k, resultC, &wg, client)
			}
		}
	}

	//listener
	go resChanListener(resultC, &wg, doneC)

	//wait for response completed
	wg.Wait()

	//if all executed, close listener
	doneC <- true
}

func ExecuteAll(project *models.Project) {
	var resultC chan *models.ResultChan = make(chan *models.ResultChan)
	var doneC chan bool = make(chan bool)

	//create waitGroup
	wg := sync.WaitGroup{}

	//create http client
	client := &http.Client{}

	//validate reqs
	if len(project.Reqs) == 0 {
		fmt.Println("No reqs found")
		close(resultC)
		close(doneC)
		return
	}

	//execute every req
	for _, v := range project.Reqs {
		wg.Add(1)
		go hit(v, resultC, &wg, client)
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
