package commands

import (
	"fmt"
)

func help(commandList map[string]CliCommand, args ...string) error {
	var formatted string

	if len(args) == 0 {
		formatted += fmt.Sprintf("Welcome to the Pokedex!\nUsage:\n")

		for _, command := range commandList {
			formatted += fmt.Sprintf("%v: %v\n", command.Name, command.Description)
		}
	} else {
		command, ok := commandList[args[0]]

		if !ok {
			printNoSuchCommand()

			return nil
		}

		formatted = fmt.Sprintf("%v\n", command.Description)
	}

	fmt.Printf("\n%v\n", formatted)

	return nil
}

func makeHelp(commandList map[string]CliCommand) func(...string) error {
	return func(args ...string) error {
		help(commandList, args...)

		return nil
	}
}
