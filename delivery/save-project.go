package delivery

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hdkef/jameter/models"
)

func SaveProject(project *models.Project) (menu int) {
	file, err := json.MarshalIndent(*project, "", " ")
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	err = os.WriteFile(fmt.Sprintf("data/%s.json", project.Name), file, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}

	fmt.Print("Project is saved")
	return -1
}
