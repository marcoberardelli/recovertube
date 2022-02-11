package service

import (
	"context"
	"log"
	"os"
	m "recovertube/model"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var ytService *youtube.Service

func init() {
	var err error

	apiyKey := os.Getenv("YOUTUBE_API_KEY")
	if apiyKey == "" {
		log.Fatalf("Missing YOUTUBE_API_KEY")
	}
	ytService, err = youtube.NewService(context.TODO(), option.WithAPIKey(apiyKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}
}

func GetVideo(id string) (m.Video, error) {

	part := []string{"snippet"}
	call := ytService.Videos.List(part).Id(id)
	result, err := call.Do()
	if err != nil {
		return m.Video{}, err
	}

	v := result.Items[0]
	//TODO: check if available
	return m.Video{
		ID:         v.Id,
		Title:      v.Snippet.Title,
		Channel:    v.Snippet.ChannelId,
		ImagePath:  v.Snippet.Thumbnails.Default.Url,
		Available:  true,
		LastUpdate: time.Now(),
	}, nil
}

func GetPlaylist(id string) (m.Playlist, error) {
	//part := []string{"snippet"}

}

func GetUserPlaylist(id string) {

}
