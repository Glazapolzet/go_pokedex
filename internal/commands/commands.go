package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/pokedex"
	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(args ...string) error
}

type CliCommands struct {
	commandList map[string]*cliCommand
	repository  repository.Repository
}

func NewCliCommands(pokedex pokedex.Pokedex, repo repository.Repository) *CliCommands {
	c := &CliCommands{
		commandList: make(map[string]*cliCommand),
		repository:  repo,
	}

	exit := makeExit()
	c.Add(exit.Name, exit)

	explore := makeExplore(repo)
	c.Add(explore.Name, explore)

	mapf := makeMapf(repo)
	c.Add(mapf.Name, mapf)

	mapb := makeMapb(repo)
	c.Add(mapb.Name, mapb)

	help := makeHelp(c.GetAll())
	c.Add(help.Name, help)

	catch := makeCatch(pokedex, repo)
	c.Add(catch.Name, catch)

	inspect := makeInspect(pokedex)
	c.Add(inspect.Name, inspect)

	pokedexCli := makePokedex(pokedex)
	c.Add(pokedexCli.Name, pokedexCli)

	return c
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
