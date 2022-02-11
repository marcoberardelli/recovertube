package model

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type YoutTubeDBRepository struct {
	db *gorm.DB
}

var ytRepo YoutTubeDBRepository

func initYouTubeRepository(dsn string) error {
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	ytRepo = YoutTubeDBRepository{db: _db}
	ytRepo.db.AutoMigrate(&User{}, &Video{}, &Playlist{})
	return nil
}

func init() {
	dsn := os.Getenv("PSQL_DSN")
	err := initYouTubeRepository(dsn)
	if err != nil {
		log.Fatalf("Error initializing YoutubeRepository")
	}
	err = initAuthRepository(dsn)
	if err != nil {
		log.Fatalf("Error initializing AuthRepository")
	}
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
	if err == nil {
		err = ErrDuplicateVideo

	} else if err == gorm.ErrRecordNotFound {
		err = r.db.Create(&video).Error
		if err != nil {
			return err
		}

		// Updating join table
		err = r.db.Table("user_video").Create(&UserVideo{userID, video.ID}).Error
	}

	return err
}

func (r YoutTubeDBRepository) SaveVideoPlaylist(video Video, userID, playlistID string) error {

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

		} else if err == ErrDuplicateVideo {
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
