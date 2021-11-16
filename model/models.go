package model

import (
	"time"
)

type YoutubeRepository interface {
	AddVideo(video Video) error
	AddPlaylist(playlist Playlist)
	GetVideo(id string) (Video, error)
	IsVideoAvailable(id string) (bool, error)
	GetPlaylist(id string) (Playlist, error)
}

type AuthRepository interface {
	Register(email, password string) error
	Login(email, password string) (User, error)
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
