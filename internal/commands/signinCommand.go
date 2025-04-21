package commands

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
)

func SignInCommand(con *config.Config, cmd Command) error {
	if len(cmd.Args) != 4 {
		return errors.New("Not enough Arguments")
	}
	username, password, err := GetUserNamePassword(cmd.Args)
	if err != nil {
		return err
	}
	ctx := context.Background()
	user, err := con.Db.GetUser(ctx, sql.NullString{String: username, Valid: true})
	if strings.Contains(err.Error(), "no rows in result") {
		return errors.New("User does not exist")
	}
	if err != nil {
		return err
	}
	if user.UserPassword.String != password {
		return errors.New("Password does not match")
	}

	// fmt.Println(newUser.UserName.String + " created")
	return nil
}
