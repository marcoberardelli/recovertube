package route

import (
	"log"
	m "recovertube/model"
	"recovertube/youtube"

	"github.com/gin-gonic/gin"
)

func GetAllVideos(c *gin.Context) {

}

func GetVideo(c *gin.Context) {

}

func AddVideo(c *gin.Context) {
	videoID := c.PostForm("video_id")
	playlistID := c.PostForm("playlist_id")

	user, ok := getUser(c)
	if !ok {
		c.JSON(500, "Internal error")
		return
	}

	video, err := youtube.GetVideo(videoID)
	if err != nil {
		return
	}

	repo, err := m.GetYTRepository()
	if err != nil {
		log.Printf("Error adding a video %+v", err)
		return
	}

	err = repo.SaveVideo(video, playlistID, user.ID)
	//TODO: fix error check
	if err != nil {
		return
	}
	//TODO: store thumbnail

	log.Printf("Added video: %s", videoID)
	c.JSON(200, "ok")

}
