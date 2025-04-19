package main

import (
	"fmt"
	"log"
	"os"

	"github.com/seidigapbar/beatdlp/db"
	"github.com/seidigapbar/beatdlp/downloader"
	"github.com/seidigapbar/beatdlp/model"
)

func main() {
	dbPath := "beatdlp.db"
	sqliteDB, err := db.NewDB(dbPath)
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	err = db.InitDB(sqliteDB)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	defer sqliteDB.Close()

	ytSvc, err := downloader.NewYoutubeClient(os.Getenv("YOUTUBE_API_KEY"))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	beatmaker := model.Beatmaker{
		Id: "erlax",
		Name: "Erlax",
		Url: "UCkCsd5HkmqYZEMm6VmBd-Yg",
	}

	db.InsertBeatmaker(sqliteDB, &beatmaker)

	instrumentals, err := downloader.Fetch(beatmaker, ytSvc)
	if err != nil {
		log.Fatalf("Error fetching instrumentals: %v", err)
	}

	for _, instrumental := range instrumentals {
		fmt.Printf("Inserting instrumental: %v\n", instrumental)
		db.InsertInstrumental(sqliteDB, &instrumental)
	}

	var dbInstrumentals []*model.Instrumental
	dbInstrumentals, err = db.GetInstrumentals(sqliteDB)
	if err != nil {
		log.Fatalf("Error getting instrumentals: %v", err)
	}

	for _, instrumental := range dbInstrumentals {
		log.Printf("Instrumental: %v", instrumental)
	}
}

