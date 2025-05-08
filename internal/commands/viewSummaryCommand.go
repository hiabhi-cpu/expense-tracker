package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
	"github.com/hiabhi-cpu/expense-tracker/internal/database"
)

func SummaryCommandDivide(con *config.Config, cmd Command) error {
	if con.User.UserName.String == "" {
		return errors.New("Please singin before adding anything")
	}
	if len(cmd.Args) == 0 {
		return SummaryCommand(con, cmd)
	} else if len(cmd.Args) == 2 {
		return MonthSummaryCommand(con, cmd)
	}
	return nil
}

func SummaryCommand(con *config.Config, cmd Command) error {
	if con.User.UserName.String == "" {
		return errors.New("Please singin before adding anything")
	}
	if len(cmd.Args) != 0 {
		return errors.New("Don't give arguments")
	}
	ctx := context.Background()

	amt_returned, err := con.Db.ViewSummary(ctx, sql.NullInt32{Int32: int32(con.User.UserID), Valid: true})
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.New("No expenses present")
		}
		// fmt.Println(err)
		return err
	}
	fmt.Println("Your total expenses are : ", amt_returned)
	return nil
}

func MonthSummaryCommand(con *config.Config, cmd Command) error {
	if con.User.UserName.String == "" {
		return errors.New("Please singin before adding anything")
	}
	if len(cmd.Args) != 2 {
		return errors.New("Don't give arguments")
	}

	month, err := GetMonth(cmd.Args)
	if err != nil {
		return err
	}
	monthTime, err := MonthNameToTime(month)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	amt_returned, err := con.Db.ViewSummaryMonth(ctx, database.ViewSummaryMonthParams{
		UserID:  sql.NullInt32{Int32: int32(con.User.UserID), Valid: true},
		Column2: monthTime,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Scan error on column index 0") {
			return errors.New(fmt.Sprint("No expenses for ", month, " present"))
		}
		// fmt.Println(err)
		return err
	}
	fmt.Println("Your total expenses for month ", month, " are : ", amt_returned)
	return nil
}

func GetMonth(args []string) (month string, err error) {
	j := 0
	for i, str := range args {
		if str == "--month" && i == 0 {
			month = args[i+1]
			j++
		}
	}
	if j != 1 {
		err = errors.New("Pass Arguments correctly")
	}
	return
}

func MonthNameToTime(monthName string) (time.Time, error) {
	monthName = strings.Title(strings.ToLower(monthName)) // "may" → "May"
	layout := "Jan"
	parsedTime, err := time.Parse(layout, monthName)
	if err != nil {
		return time.Time{}, err
	}

	// Get month from parsed time, use current year
	now := time.Now()
	return time.Date(now.Year(), parsedTime.Month(), 1, 0, 0, 0, 0, time.UTC), nil
}
