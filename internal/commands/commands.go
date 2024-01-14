package commands

import (
	"fmt"
	"os"

	"github.com/Glazapolzet/go_pokedex/internal/api"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func Get() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message.",
			Callback:    Help,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex.",
			Callback:    Exit,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations.",
			Callback:    Map,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			Callback:    Mapb,
		},
	}
}

func Help() error {
	var formatted string

	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n")

	for _, command := range Get() {
		formatted += fmt.Sprintf("%v: %v\n", command.Name, command.Description)
	}

	fmt.Printf("\n%v\n", formatted)

	return nil
}

func Exit() error {
	fmt.Printf("\nBye-bye!\n\n")

	os.Exit(0)

	return nil
}

func printMap(results []fmt.Stringer) error {
	var formatted string

	for _, locationArea := range results {
		formatted += fmt.Sprintf("%v\n", locationArea)
	}

	fmt.Printf("\n%v\n", formatted)

	return nil
}

func Map() error {
	locationAreaList := api.GetNextLocationAreaList()

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

func Mapb() error {
	locationAreaList := api.GetPrevLocationAreaList()

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
