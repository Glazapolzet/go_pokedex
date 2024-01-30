package commands

import (
	"fmt"
	"math/rand"

	"github.com/Glazapolzet/go_pokedex/internal/pokedex"
	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

func makeCatch(p pokedex.Pokedex, r repository.Repository) *cliCommand {
	callback := func(args ...string) error {
		name := args[0]

		fmt.Printf("\nThrowing a Pokeball at %v...\n", name)

		pokemon := r.GetPokemon(name)
		isCatchSuccess := checkIfSuccess(getRandomNumber(pokemon.BaseExperience))

		var formatted string

		if !isCatchSuccess {
			formatted += fmt.Sprintf("%v escaped!\n", name)
			fmt.Printf("\n%v\n", formatted)

			return nil
		}

		p.Add(pokemon)

		formatted += fmt.Sprintf("%v was caught!\n", name)
		fmt.Printf("\n%v\n", formatted)

		return nil
	}

	return &cliCommand{
		Name:        "catch",
		Description: "Catch a sertain pokemon by its name.",
		Callback:    callback,
	}
}

func checkIfSuccess(num int) bool {
	return num > 50
}

func getRandomNumber(baseExp int) int {
	return int(float64(rand.Intn(10)*10) / (float64(baseExp) / 1000))
}
