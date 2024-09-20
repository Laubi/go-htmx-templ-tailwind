package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func ConnectToDb() error {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Create users table if not exists
	err = createUserTable()
	if err != nil {
		return fmt.Errorf("failed to create user table: %w", err)
	}

	return nil
}

func CloseDb() error {
	if db == nil {
		return nil
	}

	return db.Close()
}
