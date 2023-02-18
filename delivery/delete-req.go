package delivery

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hdkef/jameter/models"
)

func DeleteRequest(project *models.Project) (menu int) {
	//show all requests
	ReadRequest(project)

	//input ids to be deleted
	fmt.Print("\nInput request ids (separated by space) :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	idStr := scanner.Text()

	idSlice := []int{}

	for _, v := range strings.Split(idStr, " ") {
		id, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid id")
			return -1
		}
		idSlice = append(idSlice, id)
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
