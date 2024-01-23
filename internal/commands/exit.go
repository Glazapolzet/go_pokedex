package commands

import (
	"fmt"
	"os"
)

func makeExit() *cliCommand {
	callback := func(args ...string) error {
		fmt.Printf("\nBye-bye!\n\n")

		os.Exit(0)

		return nil
	}

	return &cliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex.",
		Callback:    callback,
	}
}
