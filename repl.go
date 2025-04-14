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

	helloCommand := commands.Command{
		Name:        "hello",
		CommandFunc: commands.HelloCommand,
	}

	cmds.Register("hello", helloCommand)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		// parameters := words[1:]
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
