package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/pokedex"
)

func makeInspect(p pokedex.Pokedex) *cliCommand {
	callback := func(args ...string) error {
		name := args[0]

		var formatted string

		pokemon := p.Get(name)

		if pokemon == nil {
			formatted += fmt.Sprintf("You haven't caught this pokemon yet.\n")
			fmt.Printf("\n%v\n", formatted)

			return nil
		}

		formatted += fmt.Sprintf("Name: %v\n", pokemon.Name)
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
