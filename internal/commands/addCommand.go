package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
	"github.com/hiabhi-cpu/expense-tracker/internal/database"
)

func AddMoneyCommand(con *config.Config, cmd Command) error {
	// fmt.Println(con.Name)
	if con.User.UserName.String == "" {
		return errors.New("Please singin before adding anything")
	}
	if len(cmd.Args) != 4 {
		return errors.New("Not enough arguments")
	}
	desc, amt, err := GetMoneyDesc(cmd.Args)
	if err != nil {
		return err
	}
	amtInt, err := strconv.Atoi(amt)
	if err != nil {
		return errors.New("Enter correct amount in integer")
	}
	if amtInt < 0 {
		return errors.New("Amount should not be negative")
	}
	ctx := context.Background()
	newMoney, err := con.Db.CreateMoney(ctx, database.CreateMoneyParams{
		MonDesc: sql.NullString{String: desc, Valid: true},
		Amt:     sql.NullInt32{Int32: int32(amtInt), Valid: true},
		UserID:  sql.NullInt32{Int32: int32(con.User.UserID), Valid: true},
		MonDate: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		fmt.Println(err)
		return errors.New("Unable to add money")
	}
	fmt.Println(newMoney.MonID, "\t", newMoney.MonDesc, "\t", newMoney.Amt)
	fmt.Println("Added money")
	return nil
}

func GetMoneyDesc(args []string) (desc, amt string, err error) {
	j := 0
	for i, str := range args {
		if str == "--desc" && i == 0 {
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
