package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hiabhi-cpu/expense-tracker/internal/config"
	"github.com/hiabhi-cpu/expense-tracker/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

	dburl := os.Getenv("POSTGRES_URL")
	fmt.Println(dburl)

	db, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	con := &config.Config{Name: "abhi", Db: dbQueries}
	repl(con)
}
