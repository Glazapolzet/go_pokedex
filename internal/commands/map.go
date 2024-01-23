package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

func makeMapf(r repository.Repository) *cliCommand {
	callback := func(args ...string) error {
		locationAreaList := r.GetNextLocationAreaList()

		if locationAreaList == nil {
			fmt.Printf("\nNo more pages left..\n\n")

			return nil
		}

		var formatted string

		for _, locationArea := range locationAreaList.Results {
			formatted += fmt.Sprintf("%v\n", locationArea.Name)
		}

		fmt.Printf("\n%v\n", formatted)

		return nil
	}

	return &cliCommand{
		Name:        "map",
		Description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations.",
		Callback:    callback,
	}
}

func makeMapb(r repository.Repository) *cliCommand {
	callback := func(args ...string) error {
		locationAreaList := r.GetPrevLocationAreaList()

		if locationAreaList == nil {
			fmt.Printf("\nNo more pages left..\n\n")

			return nil
		}

		var formatted string

		for _, locationArea := range locationAreaList.Results {
			formatted += fmt.Sprintf("%v\n", locationArea.Name)
		}

		fmt.Printf("\n%v\n", formatted)

		return nil
	}

	return &cliCommand{
		Name:        "mapb",
		Description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
		Callback:    callback,
	}
}
