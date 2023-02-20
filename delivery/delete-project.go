package delivery

import (
	"fmt"
	"os"

	"github.com/hdkef/jameter/models"
)

func DeleteProject(project *models.Project) (menu int) {
	//input name for confirmation
	fmt.Printf("\nInput project name for confirmation : ")
	var name string
	fmt.Scanln(&name)

	if name != project.Name {
		fmt.Println("name is not match")
		return -1
	}

	//remove file
	err := os.Remove(fmt.Sprintf("data/%s.json", project.Name))
	if err != nil {
		fmt.Println("Error deleting .json file")
		return 0
	}

	fmt.Println("Project deleted")

	return 0
}
