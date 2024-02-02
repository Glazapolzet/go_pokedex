package main

import (
	"time"

	"github.com/Glazapolzet/go_pokedex/internal/commands"
	"github.com/Glazapolzet/go_pokedex/internal/pokedex/implementation"
	"github.com/Glazapolzet/go_pokedex/internal/repl"
	apiimplementation "github.com/Glazapolzet/go_pokedex/internal/repository/api_implementation"
)

var urlConfig = &apiimplementation.PokeApiUrls{
	Base:         "https://pokeapi.co/api/v2/",
	LocationArea: "https://pokeapi.co/api/v2/location-area/",
	Pokemon:      "https://pokeapi.co/api/v2/pokemon/",
}

func main() {
	pokedex := implementation.NewPokedex()
	repository := apiimplementation.NewRepo(urlConfig, time.Minute*5)
	commands := commands.NewCliCommands(pokedex, repository)
	repl := repl.NewRepl(commands)

	repl.Run()
}
