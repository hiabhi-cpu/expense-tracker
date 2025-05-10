
# Expense Tracker CLI 🧾

A simple command-line application for tracking expenses per user with support for user authentication, monthly summaries, and CRUD operations — built in Go with PostgreSQL.

https://roadmap.sh/projects/expense-tracker

## Features

- 📝 Add, update, delete expenses
- 👤 User signup, signin, signout
- 📊 View monthly summaries
- 🔐 Secure password hashing (bcrypt)
- 🗂️ PostgreSQL + SQLC based DB access


## 🚀 Getting Started

### 🛠️ Prerequisites

- [Go](https://golang.org/dl/) (v1.20+ recommended)
- [PostgreSQL](https://www.postgresql.org/)
- [`goose`](https://github.com/pressly/goose) (for DB migrations)
- `sqlc` (for generating DB code)

```bash
git clone https://github.com/hiabhi-cpu/expense-tracker.git
cd expense-tracker

# Run migrations
goose -dir db/migrations postgres "<your_connection_string>" up

# Generate SQL code
sqlc generate

# Run the CLI
go build && ./expense-tracker
```
Update the DB connection config in internal/config/config.go.

## 🧑‍💻 CLI Commands

| Command   | Description                | Usage Example                          |
|-----------|----------------------------|----------------------------------------|
| `view`    | View all past expenses     | `view`                                 |
| `summary` | View summary for a month   | `summary --month May`                  |
| `add`     | Add an expense             | `add --desc Coffee --amt 150`          |
| `update`  | Update an expense by ID    | `update --id 3 --amt 500`              |
| `delete`  | Delete an expense by ID    | `delete --id 5`                        |
| `signup`  | Register a new user        | `signup --name abhi --password 123`    |
| `signin`  | Log in                     | `signin --name abhi --password 123`    |
| `signout` | Log out                    | `signout`                              |
| `hello`   | Greet user                 | `hello`                                |
