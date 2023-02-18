package delivery

import (
	"fmt"

	"github.com/hdkef/jameter/models"
)

func MainMenu(project *models.Project) {
	var menu int = -1
	for menu != 0 {
		fmt.Print("\n\n1.Read request\n2.Create request\n3.Update request\n4.Delete request\n5.Execute request\n6.Save\n7.Quit\n8.Delete project\n\nChoose menu : ")
		_, err := fmt.Scanln(&menu)
		if err != nil {
			fmt.Println("Invalid menu")
			return
		}
		switch menu {
		case 1:
			menu = ReadRequest(project)
		case 2:
			menu = CreateRequest(project)
		case 3:
			menu = UpdateRequest(project)
		case 4:
			menu = DeleteRequest(project)
		case 5:
			menu = ExecuteRequest(project)
		case 6:
			menu = SaveProject(project)
		case 7:
			menu = 0
		case 8:
			menu = DeleteProject(project)
		}
	}
}
