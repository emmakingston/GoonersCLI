package menu

import (
	"fmt"
)

// Option represents the name of the Option and the function used to complete it
type Option struct {
	description string
	handler     func()
}

// MainOptions contains all options (description + handling func) available from the menu
var MainOptions = [...]Option{
	{"check latest score", displayLatestScore},
	{"check upcoming opponent", nil},
	{"check time until next match", nil},
}

// Display outputs each menu option to console, along with its option index
func Display() {
	for index, option := range MainOptions {
		fmt.Printf("\t%d - %s \n", index, option.description)
	}
	fmt.Println("Press any other key to exit.")
}
