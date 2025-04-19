package model

import "time"

type Instrumental struct {
	Id           string    `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	BeatmakerId  string    `json:"beatmaker_id" db:"beatmaker_id"`
	Url          string    `json:"url" db:"url"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	DownloadedAt time.Time `json:"downloaded_at" db:"downloaded_at"`
}
