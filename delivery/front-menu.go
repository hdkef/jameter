package delivery

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hdkef/jameter/models"
	"github.com/hdkef/jameter/usecase"
)

func inputProjectName(name *string) (msg string, valid bool) {
	fmt.Print("Input project name: ")
	//input project name
	_, err := fmt.Scanln(name)
	if err != nil {
		msg = "invalid input"
		return
	}

	//validate project name
	msg, valid = usecase.OnlyWordsOrDigit(*name)
	if !valid {
		return
	}
	valid = true
	return
}

func createProject(project *models.Project) (menu int, valid bool) {
	//input project name
	var name string
	msg, v := inputProjectName(&name)
	if !v {
		menu = -1
		fmt.Println(msg)
		return
	}

	//assign name to project variable
	project.Name = name
	valid = true
	return
}

func openProject(project *models.Project) (menu int, valid bool) {
	//input project name
	var name string
	msg, v := inputProjectName(&name)
	if !v {
		menu = -1
		fmt.Println(msg)
		return
	}

	//open project name .json
	byteValue, err := os.ReadFile(fmt.Sprintf("data/%s.json", name))
	if err != nil {
		menu = -1
		fmt.Println("cannot load .json")
		return
	}

	//unmarshall
	err = json.Unmarshal(byteValue, project)
	if err != nil {
		menu = -1
		fmt.Println("cannot unmarshall json")
		return
	}

	valid = true
	return
}

func FrontMenu(project *models.Project) (msg string, valid bool) {
	//msg
	var menu int = -1

	for menu != 0 {
		fmt.Print("1.Open project\n2.Create new project\n3.Exit\n\nChoose menu : ")
		_, err := fmt.Scanln(&menu)
		if err != nil {
			msg = "invalid menu"
			return
		}
		switch menu {
		case 1:
			menu, valid = openProject(project)
		case 2:
			menu, valid = createProject(project)
		case 3:
			return
		}
	}
	return
}
