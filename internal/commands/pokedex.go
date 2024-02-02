package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/pokedex"
)

func makePokedex(p pokedex.Pokedex) *cliCommand {
	callback := func(args ...string) error {
		var formatted string

		formatted += fmt.Sprintf("Your Pokedex:\n")

		for _, pokemon := range p.GetAll() {
			formatted += fmt.Sprintf(" - %v\n", pokemon.Name)
		}

		fmt.Printf("\n%v\n", formatted)

		return nil
	}

	return &cliCommand{
		Name:        "pokedex",
		Description: "Show all catched pokemons.",
		Callback:    callback,
	}
}
