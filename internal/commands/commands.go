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
	c.SetCommand(exit.Name, exit)

	explore := makeExplore(repo)
	c.SetCommand(explore.Name, explore)

	mapf := makeMapf(repo)
	c.SetCommand(mapf.Name, mapf)

	mapb := makeMapb(repo)
	c.SetCommand(mapb.Name, mapb)

	help := makeHelp(c.GetCommandList())
	c.SetCommand(help.Name, help)

	catch := makeCatch(pokedex, repo)
	c.SetCommand(catch.Name, catch)

	return c
}

func (c *CliCommands) GetCommandList() map[string]*cliCommand {
	return c.commandList
}

func (c *CliCommands) GetCommand(key string) *cliCommand {
	command, ok := c.commandList[key]

	var formatted string

	if !ok {
		formatted += fmt.Sprintf("No such command\n")
		fmt.Printf("\n%v\n", formatted)

		return nil
	}

	return command
}

func (c *CliCommands) SetCommand(key string, newCommand *cliCommand) {
	c.commandList[key] = newCommand
}
