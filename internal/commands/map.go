package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

func printEnd() {
	fmt.Printf("\nNo more pages left..\n\n")
}

func mapf(args ...string) error {
	locationAreaList := repository.GetNextLocationAreaList()

	if locationAreaList == nil {
		printEnd()

		return nil
	}

	var formatted string

	for _, locationArea := range locationAreaList.Results {
		formatted += fmt.Sprintf("%v\n", locationArea)
	}

	fmt.Printf("\n%v\n", formatted)

	return nil
}

func mapb(args ...string) error {
	locationAreaList := repository.GetPrevLocationAreaList()

	if locationAreaList == nil {
		printEnd()

		return nil
	}

	var formatted string

	for _, locationArea := range locationAreaList.Results {
		formatted += fmt.Sprintf("%v\n", locationArea)
	}

	fmt.Printf("\n%v\n", formatted)

	return nil
}
