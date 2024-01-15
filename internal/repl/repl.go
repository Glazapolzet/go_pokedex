package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Glazapolzet/go_pokedex/internal/commands"
)

func Run() {
	commandList := commands.Get()

	for {
		fmt.Print("Pokedex > ")

		prompt := getUserPrompt()

		command, ok := commandList[prompt]
		if !ok {
			fmt.Printf("\nNo such command ;(\n\n")

			continue
		}

		command.Callback()
	}
}

func getUserPrompt() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	return scanner.Text()
}
