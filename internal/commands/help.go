package commands

import (
	"fmt"
)

func help(commandList map[string]CliCommand) error {
	var formatted string

	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n")

	for _, command := range commandList {
		formatted += fmt.Sprintf("%v: %v\n", command.Name, command.Description)
	}

	fmt.Printf("\n%v\n", formatted)

	return nil
}

func makeHelp(commandList map[string]CliCommand) func() error {
	return func() error {
		help(commandList)

		return nil
	}
}
