package model

type Video struct {
	ID         string
	Title      string
	Channel    string
	Available  bool
	ImagePath  string
	LastUpdate int64
}

type Playlist struct {
	ID      string
	OwnerID string
	Videos  []Video
}
