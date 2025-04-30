package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
)

func ViewAllMoney(con *config.Config, cmd Command) error {
	if con.User.UserName.String == "" {
		return errors.New("Please singin before adding anything")
	}
	if len(cmd.Args) != 0 {
		return errors.New("Don't give arguments")
	}
	ctx := context.Background()
	moneyList, err := con.Db.GetAllMoney(ctx, sql.NullInt32{Int32: int32(con.User.UserID), Valid: true})
	if err != nil {
		fmt.Println(err)
		return errors.New("No Records or some error occured")
	}
	fmt.Println("Money Id\tAmount\tDate\t\tDescription")
	for _, arr := range moneyList {

		fmt.Println(arr.MonID, "\t\t", arr.Amt.Int32, "\t", strings.Split(arr.MonDate.Time.String(), " ")[0], "\t", arr.MonDesc.String)
	}
	return nil
}
