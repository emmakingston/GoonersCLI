package menu

import (
	"fmt"
	"strconv"
)

// Option represents the name of the option and the function used to complete it
type Option struct {
	description string
	handler     func()
}

var menuOptions = [...]string{
	"check latest score",
	"check upcoming opponent",
	"check time until next match"}

// Display outputs each menu option to console, along with its option index
func Display() {
	for index, option := range menuOptions {
		fmt.Printf("\t%d - %s \n", index, option)
	}
	fmt.Println("Press any other key to exit.")
}

// ContainsOption returns a boolean indicating if the menu choice is valid
func ContainsOption(choice string) bool {
	option, err := strconv.Atoi(choice)
	if err != nil {
		return false
	}

	if option < 0 || option > len(menuOptions) {
		return false
	}

	return true
}
