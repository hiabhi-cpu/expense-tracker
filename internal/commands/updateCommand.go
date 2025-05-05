package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
	"github.com/hiabhi-cpu/expense-tracker/internal/database"
)

func UpdateCommand(con *config.Config, cmd Command) error {
	if con.User.UserName.String == "" {
		return errors.New("Please singin before adding anything")
	}
	if len(cmd.Args) != 4 {
		return errors.New("Not enough arguments")
	}
	id, amt, err := GetMoneyId(cmd.Args)
	if err != nil {
		return err
	}
	amtInt, err := strconv.Atoi(amt)
	if err != nil {
		return errors.New("Enter correct amount in integer")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("Enter correct amount in integer")
	}
	ctx := context.Background()
	err = con.Db.UpdateMoney(ctx, database.UpdateMoneyParams{
		MonID:  int32(idInt),
		UserID: sql.NullInt32{Int32: int32(con.User.UserID), Valid: true},
		Amt:    sql.NullInt32{Int32: int32(amtInt), Valid: true},
	})
	if err != nil {
		return errors.New("Could not update")
	}
	fmt.Println("Updated the expense")
	return nil
}

func GetMoneyId(args []string) (desc, amt string, err error) {
	j := 0
	for i, str := range args {
		if str == "--id" && i == 0 {
			desc = args[i+1]
			j++
		} else if str == "--amt" && i == 2 {
			amt = args[i+1]
			j++
		}
	}
	if j != 2 {
		err = errors.New("Pass Arguments correctly")
	}
	return
}
