package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

func printMap(results []fmt.Stringer) error {
	var formatted string

	for _, locationArea := range results {
		formatted += fmt.Sprintf("%v\n", locationArea)
	}

	fmt.Printf("\n%v\n", formatted)

	return nil
}

func mapf() error {
	locationAreaList := repository.GetNextLocationAreaList()

	if locationAreaList == nil {
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

func mapb() error {
	locationAreaList := repository.GetPrevLocationAreaList()

	if locationAreaList == nil {
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
