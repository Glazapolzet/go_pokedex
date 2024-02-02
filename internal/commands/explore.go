package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

func makeExplore(r repository.Repository) *cliCommand {
	callback := func(args ...string) error {
		name := args[0]
		locationArea := r.GetLocationArea(name)

		var formatted string

		formatted += fmt.Sprintf("Found pokemon:\n")

		for _, pokemonDetails := range locationArea.PokemonEncounters {
			formatted += fmt.Sprintf("-%v\n", pokemonDetails.Pokemon.Name)
		}

		fmt.Printf("\n%v\n", formatted)

		return nil
	}

	return &cliCommand{
		Name:        "explore",
		Description: "Displays pokemons in this area.",
		Callback:    callback,
	}
}
