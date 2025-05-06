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

func DeleteCommand(con *config.Config, cmd Command) error {
	if con.User.UserName.String == "" {
		return errors.New("Please singin before adding anything")
	}
	if len(cmd.Args) != 2 {
		return errors.New("Not enough arguments")
	}
	id, err := GetId(cmd.Args)
	if err != nil {
		return err
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("Enter correct amount in integer")
	}
	ctx := context.Background()
	dbMon, err := con.Db.GetMoneyPerId(ctx, int32(idInt))
	if err != nil {
		return errors.New("Could not get the Money per ID")
	}

	if dbMon.UserID.Int32 != con.User.UserID {
		return errors.New("U cannot delete the other user money")
	}
	err = con.Db.DeleteMoney(ctx, database.DeleteMoneyParams{
		MonID:  int32(idInt),
		UserID: sql.NullInt32{Int32: int32(con.User.UserID), Valid: true},
	})
	if err != nil {
		return errors.New("Could not delete Id does not exist")
	}
	fmt.Println("Deleted the expense")

	return nil
}

func GetId(args []string) (id string, err error) {
	j := 0
	for i, str := range args {
		if str == "--id" && i == 0 {
			id = args[i+1]
			j++
		}
	}
	if j != 1 {
		err = errors.New("Pass Arguments correctly")
	}
	return
}
