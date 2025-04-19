package downloader

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func NewYoutubeClient(apiKey string) (*youtube.Service, error) {
	ctx := context.Background()

	youtubeSvc, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return youtubeSvc, nil
}
