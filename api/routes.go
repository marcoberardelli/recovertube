package api

package api

import (
	"fmt"
	"go-recover/model"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/youtube/v3"
)

var service *youtube.Service

/*
func init() {

	youtubeApiKey := os.Getenv("YOUTUBE_KEY")

	var err error
	service, err = youtube.NewService(context.Background(), option.WithAPIKey(youtubeApiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

}
*/

func AddVideo(c *gin.Context) {
	id := c.PostForm("video_id")
	if id == "" {
		c.JSON(400, "No video id")
	}

	//TODO:
	// Call to YouTube API to get info of the video
	video := model.Video{
		ID:         "videoID",
		Title:      "Come fare un webserver in go",
		Channel:    "marcoberardelli",
		Available:  true,
		ImagePath:  "adsadsad/sadsad/images",
		LastUpdate: time.Now().Unix(),
	}
	added, err := model.AddVideo(video)
	if err != nil {
		c.JSON(500, "Error on saving the video")
	}

	c.JSON(200, fmt.Sprintf("added: %t", added))

}

func AddPlaylist(c *gin.Context) {

}

func GetPlaylist(c *gin.Context) {

}

func GetUserPlaylist(c *gin.Context) {

}

func GetVideo(c *gin.Context) {

}

func GetUserVideo(c *gin.Context) {

}
