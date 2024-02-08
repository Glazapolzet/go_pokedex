package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/repository"
	"github.com/Glazapolzet/go_pokedex/internal/utils/pokedex"
	"github.com/Glazapolzet/go_pokedex/internal/utils/pokeprinter"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(args ...string) error
}

type CliCommands struct {
	commandList map[string]*cliCommand
	Repository  repository.Repository
	Pokedex     pokedex.Pokedex
	Pokeprinter pokeprinter.Pokeprinter
}

func NewCliCommands(repo repository.Repository, pokedex pokedex.Pokedex, pokeprinter pokeprinter.Pokeprinter) *CliCommands {
	c := &CliCommands{
		commandList: make(map[string]*cliCommand),
		Repository:  repo,
		Pokedex:     pokedex,
		Pokeprinter: pokeprinter,
	}

	c.AddExitCmd()
	c.AddCatchCmd()
	c.AddExploreCmd()
	c.AddHelpCmd()
	c.AddInspectCmd()
	c.AddMapbCmd()
	c.AddMapfCmd()
	c.AddPokedexCmd()

	return c
}

func (c *CliCommands) AddExitCmd() {
	exit := makeExit()
	c.Add(exit.Name, exit)
}

func (c *CliCommands) AddExploreCmd() {
	explore := makeExplore(c.Repository)
	c.Add(explore.Name, explore)
}

func (c *CliCommands) AddMapfCmd() {
	mapf := makeMapf(c.Repository)
	c.Add(mapf.Name, mapf)
}

func (c *CliCommands) AddMapbCmd() {
	mapb := makeMapb(c.Repository)
	c.Add(mapb.Name, mapb)
}

func (c *CliCommands) AddHelpCmd() {
	help := makeHelp(c.GetAll())
	c.Add(help.Name, help)
}

func (c *CliCommands) AddCatchCmd() {
	catch := makeCatch(c.Pokedex, c.Repository)
	c.Add(catch.Name, catch)
}

func (c *CliCommands) AddInspectCmd() {
	inspect := makeInspect(c.Pokedex, c.Pokeprinter)
	c.Add(inspect.Name, inspect)
}

func (c *CliCommands) AddPokedexCmd() {
	pokedexCli := makePokedex(c.Pokedex)
	c.Add(pokedexCli.Name, pokedexCli)
}

func (c *CliCommands) GetAll() map[string]*cliCommand {
	return c.commandList
}

func (c *CliCommands) Get(key string) *cliCommand {
	command, ok := c.commandList[key]

	var formatted string

	if !ok {
		formatted += fmt.Sprintf("No such command\n")
		fmt.Printf("\n%v\n", formatted)

		return nil
	}

	return command
}

func (c *CliCommands) Add(key string, newCommand *cliCommand) {
	c.commandList[key] = newCommand
}
