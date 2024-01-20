package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

func explore(args ...string) error {
	locationArea := repository.GetLocationArea(args[0])

	var formatted string

	fmt.Printf("\nFound pokemon:\n")

	for _, pokemonDetails := range locationArea.PokemonEncounters {
		formatted += fmt.Sprintf("-%v\n", pokemonDetails.Pokemon.Name)
	}

	fmt.Printf("\n%v\n", formatted)

	return nil
}
