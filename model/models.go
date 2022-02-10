package model

import (
	"errors"
	"time"
)

var (
	ErrNoRepoInstance      = errors.New("no instance of the repository")
	ErrUsedEmail           = errors.New("email is already used")
	ErrNoMatchingPassword  = errors.New("the password does not match")
	ErrDatabase            = errors.New("error during the query")
	ErrDuplicateVideo      = errors.New("the video is already stored in the database")
	ErrMultipeRepoInstance = errors.New("there is an instance of the repository")
)

type YoutubeRepository interface {
	SaveVideo(video Video, user_id string) error
	SaveVideoPlaylist(video Video, userID, playlistID string) error
	AddPlaylist(playlist Playlist, user_id string)
	GetVideo(id string) (Video, error)
	IsVideoAvailable(id string) (bool, error)
	GetPlaylist(id string) (Playlist, error)
}

type AuthenticationRepository interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
	Logout(user User) error
}

type User struct {
	ID               string `gorm:"size:12"`
	Email            string
	Password         string
	YoutubeAuthToken string
	AuthToken        string
	SavedVideos      []Video `gorm:"many2many:user_video;"`
}

type Video struct {
	ID         string
	Title      string    `gorm:"not null"`
	Channel    string    `gorm:"not null"`
	Available  bool      `gorm:"not null"`
	ImagePath  string    `gorm:"not null"`
	LastUpdate time.Time `gorm:"not null"`
}

type UserVideo struct {
	UserID  string `gorm:"primaryKey"`
	VideoID string `gorm:"primaryKey"`
}

type PlaylistVideo struct {
	PlaylistID string `gorm:"primaryKey"`
	VideoID    string `gorm:"primaryKey"`
}

type Playlist struct {
	ID      string  `gorm:"size:12"`
	OwnerID string  `gorm:"size:12"`
	Videos  []Video `gorm:"many2many:playlist_video;"`
}
