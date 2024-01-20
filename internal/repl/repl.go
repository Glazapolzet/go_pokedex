package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Glazapolzet/go_pokedex/internal/commands"
)

func Run() {
	commandList := commands.Get()

	for {
		fmt.Print("Pokedex > ")

		tokens := makeTokens(getUserPrompt())

		command, ok := commands.GetCommand(commandList, tokens[0])

		if !ok {
			continue
		}

		if len(tokens) == 1 {
			command.Callback()

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
