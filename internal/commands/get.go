package commands

func Get() map[string]CliCommand {
	commandList := map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message.",
			Callback:    nil,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays pokemons in this area.",
			Callback:    explore,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex.",
			Callback:    exit,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations.",
			Callback:    mapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.",
			Callback:    mapb,
		},
	}

	help := commandList["help"]
	help.Callback = makeHelp(commandList)
	commandList["help"] = help

	return commandList
}
