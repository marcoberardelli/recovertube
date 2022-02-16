package service

import (
	"log"

	m "recovertube/model"
	"recovertube/service/youtube"
)

func SaveVideo(videoID string, userID int32) (m.Video, error) {
	video, err := youtube.GetVideo(videoID)
	if err != nil {
		return m.Video{}, err
	}

	repo, err := m.GetYTRepository()
	if err != nil {
		log.Printf("Error adding a video %+v", err)
		return m.Video{}, err
	}

	err = repo.SaveVideo(video, userID)
	//TODO: fix error check
	if err != nil {
		return m.Video{}, err
	}
	//TODO: store thumbnail

	return video, nil
}

func SaveVideoToPlaylist(videoID, playlistID string, userID int32) error {

	ytRepo, err := m.GetYTRepository()
	if err != nil {
		log.Printf("Error retreiving yt repo")
		return err
	}
	video, err := youtube.GetVideo(videoID)
	if err != nil {
		log.Printf("Error calling youtube api")
		return err
	}
	err = ytRepo.SavePlaylistVideo(video, playlistID, userID)
	if err != nil {
		log.Printf("Error saving the video: %+v", err)
		return err
	}
	return nil
}

func SavePlaylist(playlistID, userID string) error {
	return nil
}
