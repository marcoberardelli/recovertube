package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"os"
	"recovertube/model"
	"recovertube/service"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)



func AddVideo(c *gin.Context) {
	id := c.PostForm("video_id")
	if id == "" {
		c.JSON(400, "No video id")
	}
	token, _ := c.Get("User")
	tokenSerialized, ok := token.(string)
	if !ok {
		log.Printf("Error retreiving oauth token")
		c.JSON(500, "Internal error")
	}

	
	//TODO:
	video := 
	// Call to YouTube API to get info of the video
	video := model.Video{
		ID:         "videoID",
		Title:      "Come fare un webserver in go",
		Channel:    "marcoberardelli",
		Available:  true,
		ImagePath:  "adsadsad/sadsad/images",
		LastUpdate: time.Now(),
	}
	repo, err := model.GetYTRepository()
	if err != nil {
		log.Printf("Error adding a video %+v", err)
		c.JSON(500, "Error adding video")
	}
	// TODO: get userid by token auth
	err = repo.SaveVideo(video, tokenSerialized)
	//TODO: fix error check
	if err != nil {
		c.JSON(500, "Error on saving the video")
	}

	c.JSON(200, fmt.Sprintf("added vieo:%s", video.ID))
}

func AddVideoPlaylist(c *gin.Context) {

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
