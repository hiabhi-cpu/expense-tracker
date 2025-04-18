package config

import "github.com/hiabhi-cpu/expense-tracker/internal/database"

type Config struct {
	Name string
	Db   *database.Queries
}
