package main

import (
	"time"

	"github.com/Glazapolzet/go_pokedex/internal/commands"
	"github.com/Glazapolzet/go_pokedex/internal/pokedex/implementation"
	"github.com/Glazapolzet/go_pokedex/internal/repl"
	apiimplementation "github.com/Glazapolzet/go_pokedex/internal/repository/api_implementation"
)

func main() {
	repository := apiimplementation.NewRepository(time.Minute * 5)
	pokedex := implementation.NewPokedex()

	cliCommands := commands.NewCliCommands(pokedex, repository)

	repl.Run(cliCommands)
}
