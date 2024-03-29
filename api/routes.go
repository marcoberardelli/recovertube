package api

import (
	"fmt"
	"log"
	"recovertube/auth"

	"recovertube/service"

	"github.com/gin-gonic/gin"
)

func AddVideo(c *gin.Context) {
	id := c.PostForm("video_id")
	if id == "" {
		c.JSON(400, "No video id")
		return
	}
	token, _ := c.Get("User")
	tokenSerialized, ok := token.(string)
	if !ok {
		log.Printf("Error retreiving oauth token")
		c.JSON(500, "Internal error")
		return
	}

	userCred, err := auth.DecryptJWE(tokenSerialized)
	if err != nil {
		c.JSON(500, "Error reading JWE")
		return
	}
	video, err := service.SaveVideo(id, userCred.ID)
	if err != nil {
		c.JSON(500, "Internal error")
		return
	}

	c.JSON(200, fmt.Sprintf("added vieo:%s", video.ID))
}

func AddVideoPlaylist(c *gin.Context) {
	videoID := c.PostForm("video_id")
	playlistID := c.PostForm("playlist_id")

	token, _ := c.Get("User")
	tokenSerialized, ok := token.(string)
	if !ok {
		log.Printf("Error retreiving oauth token")
		c.JSON(500, "Internal error")
	}
	userCred, err := auth.DecryptJWE(tokenSerialized)
	if err != nil {
		c.JSON(500, "Error reading JWE")
		return
	}
	service.SaveVideoToPlaylist(videoID, playlistID, userCred.ID)
	log.Printf("Added video: %s", videoID)
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
