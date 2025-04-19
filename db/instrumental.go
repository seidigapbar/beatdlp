package db

import (
	"database/sql"
	"fmt"

	"github.com/seidigapbar/beatdlp/model"
)

func InsertInstrumental(db *sql.DB, instrumental *model.Instrumental) error {
	query := `
	INSERT INTO instrumentals (id, title, beatmaker_id, url, created_at, downloaded_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(query, instrumental.Id, instrumental.Title, instrumental.BeatmakerId, instrumental.Url, instrumental.CreatedAt, instrumental.DownloadedAt)
	if err != nil {
		fmt.Printf("Failed to insert instrumental: %v\n", err)
		return fmt.Errorf("failed to insert instrumental: %w", err)
	}

	return nil
}

func GetInstrumentals(db *sql.DB) ([]*model.Instrumental, error) {
	query := `
	SELECT id, title, beatmaker_id, url, created_at, downloaded_at FROM instrumentals
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get instrumentals: %w", err)
	}

	defer rows.Close()

	instrumentals := []*model.Instrumental{}

	for rows.Next() {
		var instrumental model.Instrumental
		err := rows.Scan(&instrumental.Id, &instrumental.Title, &instrumental.BeatmakerId, &instrumental.Url, &instrumental.CreatedAt, &instrumental.DownloadedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan instrumental: %w", err)
		}
		instrumentals = append(instrumentals, &instrumental)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over instrumentals: %w", err)
	}

	return instrumentals, nil
}
