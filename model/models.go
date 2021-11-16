package model

import (
	"errors"
	"time"
)

var (
	ErrUsedEmail           = errors.New("email is already used")
	ErrNoMatchingPassword  = errors.New("the password does not match")
	ErrDatabase            = errors.New("error during the query")
	ErrDuplicateVideo      = errors.New("the video is already stored in the database")
	ErrMultipeRepoInstance = errors.New("there is an instance of the repository")
)

type YoutubeRepository interface {
	AddVideo(video Video) error
	AddPlaylist(playlist Playlist)
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
	ID               string `gorm:"size:12`
	Email            string
	Password         string
	YoutubeAuthToken string
	AUthToken        string
}

type Video struct {
	ID         string
	Title      string
	Channel    string
	Available  bool
	ImagePath  string
	LastUpdate time.Time
}

type Playlist struct {
	ID      string  `gorm:"size:12`
	OwnerID string  `gorm:"size:12`
	Videos  []Video `gorm:"many2many:playlist_videos;"`
}
