package commands

import (
	"errors"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
)

func AddCommand(con *config.Config, cmd Command) error {
	// fmt.Println(con.Name)
	if con.User.UserName.String == "" {
		return errors.New("Please singin before adding anything")
	}
	return nil
}
