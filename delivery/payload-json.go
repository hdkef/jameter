package delivery

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hdkef/jameter/models"
)

func openJSONFile(dir string, jsonStr *string) (string, bool) {
	byteValue, err := os.ReadFile(dir[1 : len(dir)-1])
	if err != nil {
		return "Cannot open .json", false
	}
	*jsonStr = string(byteValue)
	return "", true
}

func InputJSONPayload(reqs *models.ReqsWrapper) {
	var menu int = -1
	var jsonStr string
	var isDir bool
	for menu != 0 {
		//input json string
		fmt.Printf("\nInput JSON\n(for .json file use \" before & after file dir) :")
		_, err := fmt.Scanln(&jsonStr)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		//check if input is directory
		if string(jsonStr[0]) == `"` {
			isDir = true
			msg, v := openJSONFile(jsonStr, &jsonStr)
			if !v {
				fmt.Println(msg)
			} else {
				menu = 0
			}
		}

		if !isDir {
			//validate json string
			v := json.Valid([]byte(jsonStr))
			if !v {
				fmt.Println("JSON invalid")
			} else {
				menu = 0
			}
		}
	}
	reqs.Payload = models.ReqsJSON{Data: jsonStr}
}
