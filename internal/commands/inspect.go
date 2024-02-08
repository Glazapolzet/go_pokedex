package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/utils/pokedex"
	"github.com/Glazapolzet/go_pokedex/internal/utils/pokeprinter"
)

func makeInspect(px pokedex.Pokedex, pr pokeprinter.Pokeprinter) *cliCommand {
	callback := func(args ...string) error {
		name := args[0]

		var formatted string

		pokemon := px.Get(name)

		if pokemon == nil {
			formatted += fmt.Sprintf("You haven't caught this pokemon yet.\n")
			fmt.Printf("\n%v\n", formatted)

			return nil
		}

		formatted += fmt.Sprintf("Name: %v\n", pokemon.Name)
		formatted += fmt.Sprintf("\n%v\n\n", pr.GetPokemonSprite(pokemon.ID))
		formatted += fmt.Sprintf("Height: %v\n", pokemon.Height)
		formatted += fmt.Sprintf("Weight: %v\n", pokemon.Weight)

		formatted += fmt.Sprintf("Stats: \n")
		for _, s := range pokemon.Stats {
			formatted += fmt.Sprintf(" - %v: %v\n", s.Stat.Name, s.BaseStat)
		}
		formatted += fmt.Sprintf("Types: \n")
		for _, t := range pokemon.Types {
			formatted += fmt.Sprintf(" - %v\n", t.Type.Name)
		}

		fmt.Printf("\n%v\n", formatted)

		return nil
	}

	return &cliCommand{
		Name:        "inspect",
		Description: "Inspect a catched pokemon.",
		Callback:    callback,
	}
}
