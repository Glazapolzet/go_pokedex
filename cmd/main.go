package main

import (
	"time"

	"github.com/Glazapolzet/go_pokedex/config"
	"github.com/Glazapolzet/go_pokedex/internal/commands"
	"github.com/Glazapolzet/go_pokedex/internal/repl"
	apiImpl "github.com/Glazapolzet/go_pokedex/internal/repository/api_implementation"
	pokedexImpl "github.com/Glazapolzet/go_pokedex/internal/utils/pokedex/implementation"
	pokeprinterImpl "github.com/Glazapolzet/go_pokedex/internal/utils/pokeprinter/implementation"
)

var urlConfig = &apiImpl.PokeApiUrls{
	Base:         "https://pokeapi.co/api/v2/",
	LocationArea: "https://pokeapi.co/api/v2/location-area/",
	Pokemon:      "https://pokeapi.co/api/v2/pokemon/",
}

func main() {
	pokedex := pokedexImpl.NewPokedex()
	repository := apiImpl.NewRepo(urlConfig, time.Minute*5)
	pokeprinter := pokeprinterImpl.NewPokeprinter(config.SpriteList)

	commands := commands.NewCliCommands(repository, pokedex, pokeprinter)

	repl := repl.NewRepl(commands)
	repl.Run()
}
