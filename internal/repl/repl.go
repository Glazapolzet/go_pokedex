package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Glazapolzet/go_pokedex/internal/commands"
)

func Run(cliCommands *commands.CliCommands) {
	for {
		fmt.Print("Pokedex > ")

		tokens := makeTokens(getUserPrompt())

		command := cliCommands.GetCommand(tokens[0])

		if command == nil {
			continue
		}

		command.Callback(tokens[1:]...)
	}
}

func makeTokens(prompt string) []string {
	return strings.Split(prompt, " ")
}

func getUserPrompt() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()
}
