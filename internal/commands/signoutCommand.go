package commands

import (
	"errors"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
)

func SignOutCommand(con *config.Config, cmd Command) error {
	// fmt.Println(con.User.UserName.String)
	if con.User.UserName.String == "" {
		return errors.New("Please Sign in before trying to log out")
	}
	return nil
}
