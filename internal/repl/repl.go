package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Glazapolzet/go_pokedex/internal/commands"
)

type repl struct {
	commands *commands.CliCommands
}

func NewRepl(commands *commands.CliCommands) *repl {
	return &repl{
		commands: commands,
	}
}

func (r *repl) Run() {
	for {
		fmt.Print("Pokedex > ")

		tokens := r.makeTokens(r.getUserPrompt())
		keyword := tokens[0]
		args := tokens[1:]

		command := r.commands.Get(keyword)

		if command == nil {
			continue
		}

		command.Callback(args...)
	}
}

func (r *repl) makeTokens(prompt string) []string {
	return strings.Split(strings.ToLower(prompt), " ")
}

func (r *repl) getUserPrompt() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()
}
