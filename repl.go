package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hiabhi-cpu/expense-tracker/internal/commands"
	"github.com/hiabhi-cpu/expense-tracker/internal/config"
)

func repl(con *config.Config) {
	reader := bufio.NewScanner(os.Stdin)

	cmds := commands.Commands{
		RegisterCommands: make(map[string]commands.Command),
	}

	cmds.Register("hello", "", "Say hello", commands.HelloCommand)
	cmds.Register("add", "--description \"<DESCRIPTION>\" --amount <AMOUNT>", "Add an expense", commands.AddCommand)

	for {
		fmt.Print("$ Tracker > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		// parameters := words[1:]
		if commandName == "help" {
			cmds.ListCommands()
			continue
		}
		err := cmds.CommandExists(commandName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		cmds.Run(con, commandName)
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	temp := strings.Split(text, " ")
	res := []string{}
	for _, str := range temp {
		if len(str) != 0 {
			res = append(res, str)
		}
	}
	return res
}
