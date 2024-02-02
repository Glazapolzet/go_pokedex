package commands

import (
	"fmt"
)

func makeHelp(commandList map[string]*cliCommand) *cliCommand {
	helpCommand := &cliCommand{
		Name:        "help",
		Description: "Displays a help message.",
		Callback:    nil,
	}

	commandList[helpCommand.Name] = helpCommand

	callback := func(args ...string) error {
		var helpText *string

		var formatted string

		if len(args) == 0 {
			helpText = formatHelpText(commandList)
		} else {
			helpText = formatHelpTextOnCommand(commandList, args[0])

			if helpText == nil {
				formatted += fmt.Sprintf("No such command\n")
				fmt.Printf("\n%v\n", formatted)

				return nil
			}
		}

		formatted += fmt.Sprintf(*helpText)
		fmt.Printf("\n%v\n", formatted)

		return nil
	}

	helpCommand.Callback = callback

	return helpCommand
}

func formatHelpText(commandList map[string]*cliCommand) *string {
	var formatted string

	formatted += fmt.Sprintf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, command := range commandList {
		formatted += fmt.Sprintf("%v: %v\n", command.Name, command.Description)
	}

	return &formatted
}

func formatHelpTextOnCommand(commandList map[string]*cliCommand, token string) *string {
	var formatted string

	command, ok := commandList[token]

	if !ok {
		return nil
	}

	formatted += fmt.Sprintf("%v\n", command.Description)

	return &formatted
}
