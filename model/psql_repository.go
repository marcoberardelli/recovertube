package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PSQLRepository struct {
	db *gorm.DB
}

var psqlRepo PSQLRepository

func InitPSQLRepository(dsn string) (PSQLRepository, error) {

	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return PSQLRepository{}, err
	}
	psqlRepo = PSQLRepository{db: _db}
	return psqlRepo, nil
}

func GetPSQLRepository() (PSQLRepository, error) {
	if psqlRepo.db == nil {
		return PSQLRepository{}, ErrMultipeRepoInstance
	}
	return psqlRepo, nil
}

func (r PSQLRepository) AddVideo(video Video) error {

	existingVideo := Video{}
	err := r.db.First(&existingVideo, video.ID).Error
	if err == nil {
		err = ErrDuplicateVideo
	} else if err == gorm.ErrRecordNotFound {
		err = r.db.Create(&video).Error
	}

	return err
}

func (r PSQLRepository) AddPlaylist(playlist Playlist) {
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

func (r PSQLRepository) GetVideo(id string) (Video, error) {

	return Video{}, nil
}

func (r PSQLRepository) IsVideoAvailable(id string) (bool, error) {
	return false, nil
}

func (r PSQLRepository) GetPlaylist(id string) (Playlist, error) {
	return Playlist{}, nil
}
