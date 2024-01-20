package commands

import "fmt"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(args ...string) error
}

func printNoSuchCommand() {
	fmt.Printf("\nNo such command\n\n")
}

func GetCommand(commandList map[string]CliCommand, key string) (*CliCommand, bool) {
	command, ok := commandList[key]

	if !ok {
		printNoSuchCommand()

		return nil, ok
	}

	return &command, ok
}
