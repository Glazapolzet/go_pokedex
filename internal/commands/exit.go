package commands

import (
	"fmt"
	"os"
)

func makeExit() *cliCommand {
	callback := func(args ...string) error {
		var formatted string

		formatted += fmt.Sprintf("Bye-bye!\n")
		fmt.Printf("\n%v\n", formatted)

		os.Exit(0)

		return nil
	}

	return &cliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex.",
		Callback:    callback,
	}
}
