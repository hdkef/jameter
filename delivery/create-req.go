package delivery

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/hdkef/jameter/models"
	"github.com/hdkef/jameter/usecase"
)

func addHeader(reqs *models.ReqsWrapper, headers *[]models.ReqsHeader) int {
	var validator usecase.Validator
	var name string
	//input header name
	fmt.Print("Header name :")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Invalid input")
		return -1
	}

	//validate header name
	msg, v := validator.OnlyWordsOrDigit(name)
	if !v {
		fmt.Println(msg)
		return -1
	}

	//input header value
	var value string
	fmt.Print("Header value :")
	_, err = fmt.Scanln(&value)
	if err != nil {
		fmt.Println("Invalid input")
		return -1
	}

	*headers = append(*headers, models.ReqsHeader{
		Name:  name,
		Value: value,
	})
	fmt.Println("header added")
	return -1
}

func inputHeaders(reqs *models.ReqsWrapper) {
	var menu int = -1
	var headers []models.ReqsHeader
	for menu != 0 {
		fmt.Print("1.Add Header\n2.Done\n\nChoose menu :")
		_, err := fmt.Scanln(&menu)
		if err != nil {
			fmt.Println("Invalid input")
			break
		}
		switch menu {
		case 1:
			menu = addHeader(reqs, &headers)
		case 2:
			if len(headers) > 0 {
				reqs.Headers = headers
			}
			menu = 0
		}
	}
}

func addCookie(reqs *models.ReqsWrapper, cookies *[]models.ReqsCookie) int {
	var name string
	//input header name
	fmt.Print("Cookie name :")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Invalid input")
		return -1
	}

	//input header value
	var value string
	fmt.Print("Header value :")
	_, err = fmt.Scanln(&value)
	if err != nil {
		fmt.Println("Invalid input")
		return -1
	}

	*cookies = append(*cookies, models.ReqsCookie{
		Name:  name,
		Value: value,
	})
	fmt.Println("header added")
	return -1
}

func inputCookies(reqs *models.ReqsWrapper) {
	var menu int = -1
	var cookies []models.ReqsCookie
	for menu != 0 {
		fmt.Print("1.Add Cookies\n2.Done\n\nChoose menu :")
		_, err := fmt.Scanln(&menu)
		if err != nil {
			fmt.Println("Invalid input")
			break
		}
		switch menu {
		case 1:
			menu = addCookie(reqs, &cookies)
		case 2:
			if len(cookies) > 0 {
				reqs.Cookies = cookies
			}
			menu = 0
		}
	}
}

func CreateRequest(project *models.Project) (menu int) {
	var reqs models.ReqsWrapper
	var isWrapperValid bool = false
	var validator usecase.Validator
	for !isWrapperValid {
		//assign ID
		reqs.ID = len(project.Reqs) + 1

		//input name
		fmt.Print("Input request name :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name := scanner.Text()

		//validate name
		if name == "" {
			fmt.Println("Cannot empty")
			continue
		}

		reqs.Name = name

		//input Method
		fmt.Print("Input Method :")
		_, err := fmt.Scanln(&reqs.Method)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		//validate method
		msg, v := validator.OnlyWords(reqs.Method)
		if !v {
			fmt.Println(msg)
			continue
		}

		//input URI
		fmt.Print("Input URI :")
		_, err = fmt.Scanln(&reqs.URI)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		//validate URI
		_, err = url.ParseRequestURI(reqs.URI)
		if err != nil {
			fmt.Println("Invalid URI", err.Error())
			continue
		}

		//input headers
		inputHeaders(&reqs)

		//input cookies
		inputCookies(&reqs)

		//input payload type
		var ptype int
		fmt.Print("1.JSON\n2.Multipart form\n\nChoose menu\n(input anything if payload is none):")
		_, err = fmt.Scanln(&ptype)
		if err != nil {
			reqs.PayloadType = "None"
		}

		if ptype == 1 {
			reqs.PayloadType = "JSON"
		} else if ptype == 2 {
			reqs.PayloadType = "Multipart"
		} else {
			reqs.PayloadType = "None"
		}

		isWrapperValid = true
	}
	//menu payload
	switch reqs.Payload {
	case "None":
		break
	case "JSON":

	}
	fmt.Println("new request created : ", reqs)
	return -1
}