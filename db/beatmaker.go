package db

import (
	"database/sql"
	"fmt"

	"github.com/seidigapbar/beatdlp/model"
)

func InsertBeatmaker(db *sql.DB, beatmaker *model.Beatmaker) error {
	query := `
	INSERT INTO beatmakers (id, name, url)
	VALUES (?, ?, ?)
	`
	_, err := db.Exec(query, beatmaker.Id, beatmaker.Name, beatmaker.Url)
	if err != nil {
		return fmt.Errorf("failed to insert beatmaker: %w", err)
	}

	return nil
}

func GetBeatmakers(db *sql.DB) ([]*model.Beatmaker, error) {
	query := `
	SELECT id, name, url FROM beatmakers
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get beatmakers: %w", err)
	}

	defer rows.Close()

	beatmakers := []*model.Beatmaker{}

	for rows.Next() {
		var beatmaker model.Beatmaker
		err := rows.Scan(&beatmaker.Id, &beatmaker.Name, &beatmaker.Url)
		if err != nil {
			return nil, fmt.Errorf("failed to scan beatmaker: %w", err)
		}
		beatmakers = append(beatmakers, &beatmaker)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over beatmakers: %w", err)
	}

	return beatmakers, nil
}
