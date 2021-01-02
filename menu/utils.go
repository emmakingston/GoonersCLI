package menu

import (
	"errors"
	"strconv"
)

// ContainsOption returns a boolean indicating if the menu choice is valid
func ContainsOption(choice string) bool {
	option, err := strconv.Atoi(choice)
	if err != nil {
		return false
	}

	if option < 0 || option >= len(MainOptions) {
		return false
	}

	return true
}

// GetOption returns the Option struct which matches the user input
func GetOption(optionStr string) (*Option, error) {
	var option *Option

	if !ContainsOption(optionStr) {
		return option, errors.New("Menu option not found")
	}

	optionNum, _ := strconv.Atoi(optionStr)
	option = &MainOptions[optionNum]

	return option, nil
}
