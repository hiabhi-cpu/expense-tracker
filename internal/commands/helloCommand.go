package commands

import (
	"fmt"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
)

func HelloCommand(con *config.Config, cmd Command) error {
	fmt.Println(con.Name)
	return nil
}
