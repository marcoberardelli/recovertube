package api

import (
	"fmt"
	"log"

	"recovertube/model"
	"recovertube/youtube"

	"github.com/gin-gonic/gin"
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

	video, err := youtube.GetVideo(id)
	if err != nil {
		log.Fatalf("Error getting video info from youtube")
		c.JSON(500, "Error getting video info from youtube")
	}
	/*
		video := model.Video{
			ID:         "videoID",
			Title:      "Come fare un webserver in go",
			Channel:    "marcoberardelli",
			Available:  true,
			ImagePath:  "adsadsad/sadsad/images",
			LastUpdate: time.Now(),
		}

	*/

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
	//TODO: add in youtube playlist
	videoID := c.PostForm("video_id")
	playlistID := c.PostForm("playlist_id")

	token, _ := c.Get("User")
	tokenSerialized, ok := token.(string)
	if !ok {
		log.Printf("Error retreiving oauth token")
		c.JSON(500, "Internal error")
	}
	ytRepo, err := model.GetYTRepository()
	if err != nil {
		log.Printf("Error retreiving yt repo")
		c.JSON(500, "Internal error ")
	}
	video, err := youtube.GetVideo(videoID)
	if err != nil {
		log.Printf("Error calling youtube api")
		c.JSON(500, "Internal error")
	}
	err = ytRepo.SaveVideoPlaylist(video, tokenSerialized, playlistID)
	if err != nil {
		log.Printf("Error saving the video: %+v", err)
		c.JSON(500, "Internal error")
	}

	log.Printf("Added video: %s", video.Title)
	c.JSON(200, "ok")

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
