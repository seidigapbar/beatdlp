package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, nil
}

func InitDB(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS beatmakers (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		url TEXT NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS instrumentals (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		beatmaker_id TEXT NOT NULL,
		url TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		downloaded_at TIMESTAMP,
		FOREIGN KEY (beatmaker_id) REFERENCES beatmakers(id)
	)
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}
