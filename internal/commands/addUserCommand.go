package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
	"github.com/hiabhi-cpu/expense-tracker/internal/database"
)

func AddUserCommand(con *config.Config, cmd Command) error {
	if len(cmd.Args) != 4 {
		return errors.New("Not enough Arguments")
	}
	username, password, err := getUserNamePassword(cmd.Args)
	if err != nil {
		return err
	}
	ctx := context.Background()
	newUser, err := con.Db.CreateUser(ctx, database.CreateUserParams{
		UserName:     sql.NullString{String: username, Valid: true},
		UserPassword: sql.NullString{String: password, Valid: true},
	})
	if strings.Contains(err.Error(), "duplicate key") {
		return errors.New("User already exist")
	}
	if err != nil {
		return err
	}

	fmt.Println(newUser.UserName.String + " created")
	return nil
}

func getUserNamePassword(args []string) (username, password string, err error) {
	j := 0
	for i, str := range args {
		if str == "--name" && i == 0 {
			username = args[i+1]
			j++
		} else if str == "--password" && i == 2 {
			password = args[i+1]
			j++
		}
	}
	if j != 2 {
		err = errors.New("Pass Arguments correctly")
	}
	return
}
