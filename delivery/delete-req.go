package delivery

import (
	"fmt"

	"github.com/hdkef/jameter/models"
	"github.com/hdkef/jameter/usecase"
)

func DeleteRequest(project *models.Project) (menu int) {
	//show all requests
	ReadRequest(project)

	//input ids to be deleted
	prompter := usecase.Prompt{}
	idSlice, msg, v := prompter.GetReqIDSlice()
	if !v {
		fmt.Println(msg)
		return -1
	}

	//delete from projects
	for _, v := range idSlice {
		for k := 0; k < len(project.Reqs); k++ {
			if v == project.Reqs[k].ID {
				//if found id to be deleted  then deleted it
				project.Reqs = append(project.Reqs[:k], project.Reqs[k+1:]...)
				break
			}
		}
	}

	fmt.Println("request deleted with ids ", idSlice)
	return -1
}
