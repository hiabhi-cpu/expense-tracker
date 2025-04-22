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
	username, password, err := GetUserNamePassword(cmd.Args)
	if err != nil {
		return err
	}
	ctx := context.Background()
	hashedPass, err := HashPassword(password)
	if err != nil {
		return errors.New("Password could not be hashed")
	}
	// fmt.Println(hashedPass)
	newUser, err := con.Db.CreateUser(ctx, database.CreateUserParams{
		UserName:     sql.NullString{String: username, Valid: true},
		UserPassword: sql.NullString{String: hashedPass, Valid: true},
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return errors.New("User already exists")
		}
		return err
	}

	fmt.Println(newUser.UserName.String + " created")
	return nil
}
