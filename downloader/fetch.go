package downloader

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/youtube/v3"

	"github.com/seidigapbar/beatdlp/model"
)

func humanizeTitle(title string) string {
	// Remove special characters and replace spaces with underscores
	cleanTitle := strings.Map(func(r rune) rune {
		if r == ' ' {
			return '_'
		}

		if (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == '_' || r == '-' {
			return r
		}

		return -1
	}, title)

	// Convert to lowercase for consistency
	return strings.ToLower(cleanTitle)
}

func Fetch(beatmaker model.Beatmaker, searchService *youtube.Service) ([]model.Instrumental, error) {
	search := searchService.Search.List([]string{"id", "snippet"}).ChannelId(beatmaker.Url).MaxResults(50).Order("date")

	searchResult, err := search.Do()

	if err != nil {
		// TODO logs here
		fmt.Println("Error fetching beatmaker data")
		fmt.Println(err)
		return nil, err
	}

	var result []model.Instrumental
	for _, items := range searchResult.Items {

		var beatmakerInstrumental model.Instrumental
		beatmakerInstrumental.Id = uuid.New().String()

		publishedAt, err := time.Parse(time.RFC3339, items.Snippet.PublishedAt)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}

		beatmakerInstrumental.CreatedAt = publishedAt
		beatmakerInstrumental.Url = items.Id.VideoId
		beatmakerInstrumental.Title = items.Snippet.Title
		beatmakerInstrumental.HumanizedTitle = humanizeTitle(items.Snippet.Title)
		beatmakerInstrumental.BeatmakerId = beatmaker.Id

		// TODO: REDO THE LOGIC HERE
		beatmakerInstrumental.DownloadedAt = time.Now()

		result = append(result, beatmakerInstrumental)
	}

	return result, nil
}
