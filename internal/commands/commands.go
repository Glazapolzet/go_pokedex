package commands

import (
	"fmt"

	"github.com/Glazapolzet/go_pokedex/internal/repository"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(args ...string) error
}

func NewCliCommands(repo repository.Repository) *CliCommands {
	c := &CliCommands{
		commandList: map[string]*cliCommand{},
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

	return c
}

type CliCommands struct {
	commandList map[string]*cliCommand
	repository  repository.Repository
}

func (c *CliCommands) GetCommandList() map[string]*cliCommand {
	return c.commandList
}

func (c *CliCommands) GetCommand(key string) *cliCommand {
	command, ok := c.commandList[key]

	if !ok {
		fmt.Printf("\nNo such command\n\n")

		return nil
	}

	return command
}

func (c *CliCommands) SetCommand(key string, newCommand *cliCommand) {
	c.commandList[key] = newCommand
}
