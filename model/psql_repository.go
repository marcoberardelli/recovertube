package model

import (
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	ErrDatabase            = errors.New("Error during the query")
	ErrNoMatchingPassword  = errors.New("The password does not match")
	ErrDuplicateVideo      = errors.New("The video is already stored in the database")
	ErrMultipeRepoInstance = errors.New("There is an instance of the repository")
)

type PSQLRepository struct {
	db *gorm.DB
}

var psqlRepo PSQLRepository

func InitPSQLRepository(dsn string) (PSQLRepository, error) {

	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return PSQLRepository{}, err
	}
	return PSQLRepository{db: _db}, nil
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
	if err == gorm.ErrRecordNotFound {
		return ErrDuplicateVideo
	} else if err != nil {
		return err
	}
	err = r.db.Create(&video).Error

	if err != nil {
		return err
	}

	return nil
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
