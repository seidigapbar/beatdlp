// TODO: yt-dlp adapter

package downloader

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/lrstanley/go-ytdlp"
	"github.com/seidigapbar/beatdlp/model"
)

func Download(instrumental model.Instrumental, ytDlp *ytdlp.Command) ([]*ytdlp.ExtractedInfo, error) {
	// Get humanized filename

	result, err := ytDlp.
		ExtractAudio().
		AudioFormat("mp3").
		Output(filepath.Join("downloads", instrumental.HumanizedTitle)).
		Run(context.Background(), instrumental.Url)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Downloaded to: %s\n", instrumental.HumanizedTitle)
	return result.GetExtractedInfo()
}
