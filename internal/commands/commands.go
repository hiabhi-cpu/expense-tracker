package commands

import (
	"errors"
	"fmt"

	"github.com/fatih/color"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
)

type Command struct {
	Name        string
	ArgsList    string
	Args        []string
	CommandFunc func(*config.Config, Command) error
}

type Commands struct {
	RegisterCommands map[string]Command
}

func (c *Commands) Register(name string, argsList string, f func(*config.Config, Command) error) {
	comm := Command{
		Name:        name,
		ArgsList:    argsList,
		CommandFunc: f,
	}
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

func (c *Commands) ListCommands() {
	fmt.Println("Commands List:")
	for i, cmd := range c.RegisterCommands {

		fmt.Println(i)
		fmt.Printf(i + " ")
		flag := 0
		for _, s := range cmd.ArgsList {
			if s == '<' {
				flag = 1
			}
			if flag == 1 {
				// color.Red.Print(string(s))
				color.New(color.FgRed).Print(string(s))
			} else {
				fmt.Print(string(s))
			}
			if s == '>' {
				flag = 0
			}
		}
		fmt.Println()
	}
}
