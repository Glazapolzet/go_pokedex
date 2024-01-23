package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Glazapolzet/go_pokedex/internal/commands"
	apiimplementation "github.com/Glazapolzet/go_pokedex/internal/repository/api_implementation"
)

func Run() {
	repository := apiimplementation.NewRepository(time.Minute * 5)
	cliCommands := commands.NewCliCommands(repository)

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
