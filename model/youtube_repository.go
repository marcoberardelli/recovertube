package model

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type YoutTubeDBRepository struct {
	db *gorm.DB
}

var ytRepo YoutTubeDBRepository

func init() {
	dsn := os.Getenv("PSQL_DSN")
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error initializing yt repo")
	}
	ytRepo = YoutTubeDBRepository{db: _db}
	ytRepo.db.AutoMigrate(&User{}, &Video{}, &Playlist{})
}

func GetYTRepository() (YoutTubeDBRepository, error) {
	if ytRepo.db == nil {
		return YoutTubeDBRepository{}, ErrNoRepoInstance
	}
	return ytRepo, nil
}

func (r YoutTubeDBRepository) SaveVideo(video Video, userID string) error {

	existingVideo := Video{}
	err := r.db.First(&existingVideo, "id = ?", video.ID).Error
	if err == gorm.ErrRecordNotFound {
		err = r.db.Create(&video).Error
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	existingVideo.LastUpdate = time.Now()
	err = r.db.Save(&existingVideo).Error
	if err != nil {
		log.Printf("Error updating video.LastUpdate attribute")
		return err
	}

	// Updating join table
	err = r.db.Table("user_video").Create(&UserVideo{userID, video.ID}).Error
	return err
}

func (r YoutTubeDBRepository) SavePlaylistVideo(video Video, playlistID, userID string) error {

	err := r.SaveVideo(video, userID)
	if err != nil {
		return err
	}

	err = r.db.Table("playlist_video").Create(&PlaylistVideo{playlistID, video.ID}).Error

	return err
}

func (r YoutTubeDBRepository) AddPlaylist(playlist Playlist, user_id string) error {
	for _, v := range playlist.Videos {
		err := r.SaveVideo(v, user_id)
		if err == nil {
			// Correctly added the video
			log.Printf("Added %s", v.Title)

		} else if err == ErrExistingVideo {
			// Already in db
			log.Printf("Already present %s : %s", v.ID, v.Title)
			// Updating join table
			res := r.db.Table("playlist_video").Create(&PlaylistVideo{playlist.ID, v.ID})
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

func (r YoutTubeDBRepository) GetVideo(id string) (Video, error) {

	// db.Omit().Query()
	return Video{}, nil
}

func (r YoutTubeDBRepository) IsVideoAvailable(id string) (bool, error) {
	return false, nil
}

func (r YoutTubeDBRepository) GetPlaylist(id string) (Playlist, error) {
	return Playlist{}, nil
}
