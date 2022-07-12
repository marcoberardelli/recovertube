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

package model

import (
	"log"
	"time"

	"gorm.io/gorm"
)

func SaveVideo(video Video, playlistID string, userID int32) error {

	existingVideo := Video{}
	err := db.First(&existingVideo, "id = ?", video.ID).Error
	if err == gorm.ErrRecordNotFound {
		err = db.Create(&video).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	existingVideo.LastUpdate = time.Now()
	err = db.Save(&existingVideo).Error
	if err != nil {
		log.Printf("Error updating video.LastUpdate attribute")
		return err
	}

	// Updating join table
	err = db.Table("playlist_video").Create(&PlaylistVideo{playlistID, video.ID}).Error
	if err != nil {
		log.Printf("Error updating join playlist_video table")
		return err
	}
	err = db.Table("user_video").Create(&UserVideo{userID, video.ID}).Error
	return err
}

func AddPlaylist(playlist Playlist, userID int32) error {
	for _, v := range playlist.Videos {
		err := SaveVideo(v, playlist.ID, userID)
		if err == nil {
			// Correctly added the video
			log.Printf("Added %s", v.Title)

		} else if err == ErrExistingVideo {
			// Already in db
			log.Printf("Already present %s : %s", v.ID, v.Title)
			// Updating join table
			res := db.Table("playlist_video").Create(&PlaylistVideo{playlist.ID, v.ID})
			if res.Error != nil {
				log.Fatalf(res.Error.Error())
			}

		} else {
			// Error on insert
			log.Printf("Error adding video ID:%s,  Title:%s   Error: %s", v.ID, v.Title, err)
			return err
		}
	}
	return nil
}

func GetVideo(id string) (Video, error) {

	// db.Omit().Query()
	return Video{}, nil
}

func IsVideoAvailable(id string) (bool, error) {
	return false, nil
}

func GetPlaylist(id string) (Playlist, error) {
	return Playlist{}, nil
}

func GetAllPlaylists(userID int32) ([]Playlist, error) {
	return []Playlist{}, nil
}
