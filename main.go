package main

import (
	"fmt"

	"github.com/hdkef/jameter/delivery"
	"github.com/hdkef/jameter/models"
)

func main() {
	//shows front menu
	var project models.Project
	msg, valid := delivery.FrontMenu(&project)
	if !valid {
		fmt.Println(msg)
		return
	}

	//shows main menu
	delivery.MainMenu(&project)

	fmt.Println("jameter by hdkef")
}
