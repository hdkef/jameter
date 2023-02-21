package delivery

import (
	"fmt"

	"github.com/hdkef/jameter/models"
)

func MainMenu(project *models.Project) {
	var menu int = -1
	for menu != 0 {
		fmt.Print("\n1.Read request\n2.Create request\n3.Delete request\n4.Execute request\n5.Load Test\n6.Save\n7.Quit\n8.Delete project\n\nChoose menu : ")
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
			menu = DeleteRequest(project)
		case 4:
			menu = ExecuteRequest(project)
		case 5:
			menu = LoadTest(project)
		case 6:
			menu = SaveProject(project)
		case 7:
			menu = 0
		case 8:
			menu = DeleteProject(project)
		}
	}
}
