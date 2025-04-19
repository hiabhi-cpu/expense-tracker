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
	CmdDesc     string
}

type Commands struct {
	RegisterCommands map[string]Command
}

func (c *Commands) Register(name string, argsList string, cmdDesc string, f func(*config.Config, Command) error) {
	comm := Command{
		Name:        name,
		ArgsList:    argsList,
		CommandFunc: f,
		CmdDesc:     cmdDesc,
	}
	c.RegisterCommands[name] = comm
}

func (c *Commands) Run(con *config.Config, cmdName string, parameters []string) error {
	cmd, e := c.RegisterCommands[cmdName]
	if !e {
		return errors.New("Command not found")
	}
	cmd.Args = parameters
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
		c.PrintCommand(i, cmd)

	}
}

func (c *Commands) PrintCommand(cmdName string, cmd Command) {
	fmt.Println("\t" + cmdName + "\t" + cmd.CmdDesc)
	fmt.Printf("\tUsage: " + cmdName + " ")
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
	fmt.Println("\n")
}
