package model

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CompareLogin(email, password string) bool {

	var _email, _password string
	row := db.QueryRow(queryBuilder.CompareLoginQuery())
	if err := row.Scan(&email, &password); err != nil {
		return false
	}
	if (email == _email) && (password == _password) {
		return true
	}

	return false
}

func AddVideo(video Video) (bool, error) {
	//TODO: search for row already in db
	added := false
	stmt, err := db.Prepare(queryBuilder.AddVideoQuery())
	if err != nil {
		log.Printf("Error creating statement for video [%s]", video.ID)
		return added, err
	}

	result, err := stmt.Exec(video.ID, video.Title, video.Channel, video.Available, video.ImagePath)
	if err != nil {
		log.Printf("Error executing statement for video [%s]", video.ID)
		return added, err
	}

	rows, err := result.RowsAffected()
	if rows >= 0 {
		added = true
	}
	if err != nil {
		return added, err
	}
	return added, nil
}

func AddPlaylist(playlist Playlist) {
	for _, v := range playlist.Videos {
		isAdded, err := AddVideo(v)
		if isAdded {
			// Correctly added the video
			log.Printf("Added %s", v.Title)
		} else if err != nil {
			// Error on insert
			log.Printf("Error adding video ID:%s,  Title:%s   Error: %s", v.ID, v.Title, err)
		} else {
			// Already in db
			log.Printf("Already present %s : %s", v.ID, v.Title)
		}
	}
}

func GetVideo(id string) (Video, error) {
	row := db.QueryRow(queryBuilder.GetVideoQuery(), id)
	var video Video
	err := row.Scan(&video.ID, &video.Title, &video.Channel, &video.Available, &video.ImagePath)
	if err != nil {
		log.Printf("Couldn't get info from video [%s]", id)
		return Video{}, err
	}

	return video, nil
}

func IsVideoAvailable(id string) (bool, error) {
	var isAvailable bool

	row := db.QueryRow(queryBuilder.IsVideoAvailableQuery(), id)
	err := row.Scan(&isAvailable)
	if err != nil {
		log.Printf("Error retreiving the status of video [%s]", id)
		return false, err
	}
	return isAvailable, nil
}

func GetPlaylist(id string) (Playlist, error) {

	var playlist Playlist
	row := db.QueryRow(queryBuilder.GetPlaylistQuery(), id)
	err := row.Scan(&playlist.ID, &playlist.OwnerID)
	if err != nil {
		log.Printf("Couldn't get playlist data [%s]: %s", id, err)
		return Playlist{}, err
	}

	rows, err := db.Query("SELECT Videos FROM Videos, Video_Playlist WHERE Videos.id=Video_Playlist.video_id AND Video_Playlist.playlist_id=?", id)
	if err != nil {
		log.Printf("Couldn't get videos from the playlist [%s]: %s", id, err)
		return Playlist{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var video Video
		err = rows.Scan(&video.ID, &video.Title, &video.Channel, &video.ImagePath)
		if err != nil {
			log.Printf("Error scanning the video result: %s", err)
		} else {
			playlist.Videos = append(playlist.Videos, video)
		}
	}

	return playlist, nil
}
