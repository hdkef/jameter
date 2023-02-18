package delivery

import (
	"encoding/json"
	"fmt"

	"github.com/hdkef/jameter/models"
)

func ReadRequest(project *models.Project) (menu int) {
	if len(project.Reqs) == 0 {
		fmt.Println("No request found")
		return -1
	}
	for _, v := range project.Reqs {
		fmt.Printf("ID\t\t: %d\n", v.ID)
		fmt.Printf("Name\t\t: %s\n", v.Name)
		fmt.Printf("Method\t\t: %s\n", v.Method)
		fmt.Printf("URI\t\t: %s\n", v.URI)
		fmt.Printf("Headers\t\t: %s\n", v.Headers)
		fmt.Printf("Cookies\t\t: %s\n", v.Cookies)
		fmt.Printf("Payload type\t: %s\n", v.PayloadType)
		if v.PayloadType == "JSON" {
			d, err := json.Marshal(v.Payload)
			if err != nil {
				fmt.Println("payload type cast fail")
				return -1
			}
			var p models.ReqsJSON
			err = json.Unmarshal(d, &p)
			if err != nil {
				fmt.Println("payload type cast fail")
				return -1
			}
			fmt.Printf("Payload\t\t: %s\n", p.Data)
		}
		fmt.Println("")
	}
	return -1
}
