package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
)

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
