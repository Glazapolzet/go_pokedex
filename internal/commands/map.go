package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

func printEnd() {
	fmt.Printf("\nNo more pages left..\n\n")
}

func printMap(results []fmt.Stringer) error {
	var formatted string

	for _, element := range results {
		formatted += fmt.Sprintf("%v\n", element)
	}

	fmt.Printf("\n%v\n", formatted)

	return nil
}

func mapf(args ...string) error {
	locationAreaList := repository.GetNextLocationAreaList()

	if locationAreaList == nil {
		printEnd()

		return nil
	}

	results := locationAreaList.Results

	casted := make([]fmt.Stringer, 0, len(results))

	for _, v := range results {
		casted = append(casted, v)
	}

	printMap(casted)

	return nil
}

func mapb(args ...string) error {
	locationAreaList := repository.GetPrevLocationAreaList()

	if locationAreaList == nil {
		printEnd()

		return nil
	}

	results := locationAreaList.Results

	casted := make([]fmt.Stringer, 0, len(results))

	for _, v := range results {
		casted = append(casted, v)
	}

	printMap(casted)

	return nil
}
