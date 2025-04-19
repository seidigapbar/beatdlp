package model

type Beatmaker struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Url    string `json:"url" db:"url"`
}