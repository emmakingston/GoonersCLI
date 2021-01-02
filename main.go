package main

import (
	"fmt"

	"github.com/emmakingston/goonerscli/menu"
)

func main() {
	fmt.Println("Welcome to the Gooners CLI")
	printMenu()
	fmt.Println("Thank you for using the Gooners CLI")
}

func printMenu() {
	var userChoice string

	for {
		fmt.Println("Please choose a service from the below: ")
		menu.Display()
		fmt.Scan(&userChoice)

		if menu.ContainsOption(userChoice) {
			menu.ProcessChoice(userChoice)
			fmt.Println()
		} else {
			break
		}
	}
}
