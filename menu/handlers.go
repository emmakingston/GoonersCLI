package menu

import (
	"errors"
	"fmt"
)

// ProcessChoice carries out the action requested by the user
func ProcessChoice(choice string) {
	option, err := GetOption(choice)
	if err != nil {
		fmt.Println(err)
		return
	}

	if option.handler == nil {
		fmt.Println(errors.New("No handling function defined"))
		return
	}

	option.handler()
}

func displayLatestScore() {
	fmt.Println("Last score was 0-0")
}
