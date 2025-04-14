package commands

import (
	"errors"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
)

type Command struct {
	Name        string
	Args        []string
	CommandFunc func(*config.Config, Command) error
}

type Commands struct {
	RegisterCommands map[string]Command
}

func (c *Commands) Register(name string, comm Command) {
	c.RegisterCommands[name] = comm
}

func (c *Commands) Run(con *config.Config, cmdName string) error {
	cmd, e := c.RegisterCommands[cmdName]
	if !e {
		return errors.New("Command not found")
	}
	return cmd.CommandFunc(con, cmd)
}

func (c *Commands) CommandExists(cmdName string) error {
	_, err := c.RegisterCommands[cmdName]
	if !err {
		return errors.New("Command not found")
	}
	return nil
}
