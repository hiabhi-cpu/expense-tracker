package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("User already exists")
		}
		return err
	}
	if !CheckPassword(user.UserPassword.String, password) {
		return errors.New("Password does not match")
	}

	fmt.Println("Hello" + user.UserName.String)
	return nil
}
