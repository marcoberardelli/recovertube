package service

import (
	"log"
	"recovertube/model"
	"recovertube/service/youtube"
)

func SaveVideo(videoID, userID string) error {
	video, err := youtube.GetVideo(videoID)
	if err != nil {
		return err
	}

	repo, err := model.GetYTRepository()
	if err != nil {
		log.Printf("Error adding a video %+v", err)
		return err
	}

	err = repo.SaveVideo(video, userID)
	//TODO: fix error check
	if err != nil {
		return err
	}

	return nil
}

func SaveVideoToPlaylist(videoID, playlistID, userID string) error {

	ytRepo, err := model.GetYTRepository()
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
