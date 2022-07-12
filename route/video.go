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

package route

import (
	"log"
	"recovertube/model"
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

	err = model.SaveVideo(video, playlistID, user.ID)
	//TODO: fix error check
	if err != nil {
		return
	}
	//TODO: store thumbnail

	log.Printf("Added video: %s", videoID)
	c.JSON(200, "ok")

}
