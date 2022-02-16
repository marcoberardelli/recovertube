package model

import (
	"errors"
	"time"
)

var (
	ErrNoRepoInstance     = errors.New("missing instance of repository")
	ErrUsedEmail          = errors.New("email is already used")
	ErrNoMatchingPassword = errors.New("the password does not match")
	ErrDatabase           = errors.New("error during the query")
	ErrExistingVideo      = errors.New("the video is already stored in the database")
	ErrUserNotRegistered  = errors.New("email is not registered")
)

type YoutubeRepository interface {
	GetVideo(id string, userID int32) (Video, error)
	GetAllVideos(id string, userID int32) ([]Video, error)
	SaveVideo(video Video, playlistID string, userID int32) error
	IsVideoAvailable(id string) (bool, error)
	GetPlaylist(id string, userID int32) (Playlist, error)
	GetAllPlaylists(userID int32) ([]Playlist, error)
	AddPlaylist(playlist Playlist, userID int32) error
	NewPlaylist(playlist Playlist, userID int32) error
}

type ThumbnailStore interface {
	SaveImage(url string)
	GetImage(videoID string)
}

type User struct {
	ID    int32
	Email string
	//Playlists   []Playlist
	SavedVideos []Video `gorm:"many2many:user_video;"`
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
	UserID  int32  `gorm:"primaryKey"`
	VideoID string `gorm:"primaryKey"`
}

type PlaylistVideo struct {
	PlaylistID string `gorm:"primaryKey"`
	VideoID    string `gorm:"primaryKey"`
}

type Playlist struct {
	ID      string  `gorm:"size:12"`
	OwnerID int32   `gorm:"size:12"`
	Owner   User    `gorm:"size:12;foreignKey:OwnerID"`
	Videos  []Video `gorm:"many2many:playlist_video;"`
}
