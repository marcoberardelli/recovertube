// Copyright (C) 2022  Marco Berardelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package youtube

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
	return m.Playlist{}, nil
}

func GetUserPlaylist(id string) {

}
