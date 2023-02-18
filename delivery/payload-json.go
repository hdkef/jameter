package delivery

import (
	"encoding/json"
	"fmt"

	"github.com/hdkef/jameter/models"
)

func InputJSONPayload(reqs *models.ReqsWrapper) {
	var menu int = -1
	var jsonStr string
	for menu != 0 {
		//input json string
		fmt.Printf("Input JSON :")
		_, err := fmt.Scanln(&jsonStr)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		//validate json string
		valid := json.Valid([]byte(jsonStr))
		if !valid {
			fmt.Println("JSON invalid")
		} else {
			menu = 0
		}
	}
	reqs.Payload = models.ReqsJSON{Data: jsonStr}
}
