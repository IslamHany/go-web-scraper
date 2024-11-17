package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "../api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createProductsTable()
}

func createProductsTable() {
	createProductsTable := `
		CREATE TABLE IF NOT EXISTS products(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createProductsTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create products table")
	}
}
