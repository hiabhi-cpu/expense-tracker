package commands

import (
	"errors"
	"fmt"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
	"github.com/hiabhi-cpu/expense-tracker/internal/database"
)

func SignOutCommand(con *config.Config, cmd Command) error {
	// fmt.Println(con.User.UserName.String)
	if con.User.UserName.String == "" {
		return errors.New("Please Sign in before trying to signout out")
	}
	con.User = database.User{}
	fmt.Println("Signed Out Successfully")
	return nil
}
