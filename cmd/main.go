package main

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/commands"
	"github.com/Glazapolzet/go_pokedex/internal/repl"
)

func main() {
	commandList := commands.Get()

	for {
		fmt.Print("Pokedex > ")

		prompt := repl.GetUserPrompt()

		command, ok := commandList[prompt]

		if !ok {
			fmt.Printf("\nNo such command ;(\n\n")
			continue
		}

		command.Callback()
	}
}
