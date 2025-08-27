package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init(path string) error {
	var err error
	DB, err = sql.Open("sqlite", path)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Create a table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		ticket INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT
	);
	`
	_, err = DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}
