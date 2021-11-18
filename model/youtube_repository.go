package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type YoutTubeDBRepository struct {
	db *gorm.DB
}

var ytRepo YoutTubeDBRepository

func InitYTRepository(dsn string) (YoutTubeDBRepository, error) {

	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return YoutTubeDBRepository{}, err
	}
	ytRepo = YoutTubeDBRepository{db: _db}
	ytRepo.db.AutoMigrate(&User{}, &Video{}, &Playlist{})
	return ytRepo, nil
}

func GetYTRepository() (YoutTubeDBRepository, error) {
	if ytRepo.db == nil {
		return YoutTubeDBRepository{}, ErrNoRepoInstance
	}
	return ytRepo, nil
}

func (r YoutTubeDBRepository) AddVideo(video Video) error {

	existingVideo := Video{}
	err := r.db.First(&existingVideo, video.ID).Error
	if err == nil {
		err = ErrDuplicateVideo
	} else if err == gorm.ErrRecordNotFound {
		err = r.db.Create(&video).Error
	}

	return err
}

func (r YoutTubeDBRepository) AddPlaylist(playlist Playlist) {
	for _, v := range playlist.Videos {
		err := r.AddVideo(v)
		if err == nil {
			// Correctly added the video
			log.Printf("Added %s", v.Title)
		} else if err == ErrDuplicateVideo {
			// Already in db
			log.Printf("Already present %s : %s", v.ID, v.Title)
		} else {
			// Error on insert
			log.Printf("Error adding video ID:%s,  Title:%s   Error: %s", v.ID, v.Title, err)
		}
	}
}

func (r YoutTubeDBRepository) GetVideo(id string) (Video, error) {

	return Video{}, nil
}

func (r YoutTubeDBRepository) IsVideoAvailable(id string) (bool, error) {
	return false, nil
}

func (r YoutTubeDBRepository) GetPlaylist(id string) (Playlist, error) {
	return Playlist{}, nil
}
