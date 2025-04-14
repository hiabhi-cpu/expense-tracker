package main

import "github.com/hiabhi-cpu/expense-tracker/internal/config"

func main() {
	con := &config.Config{Name: "abhi"}
	repl(con)
}
