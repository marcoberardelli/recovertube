package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psqlDriver := os.Getenv("POSTGRESQL_DRIVER")

	service, err := youtube.NewService(context.Background(), option.WithAPIKey(YOUTUBE_KEY))
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	service.Playlists.Insert()
	/*
		r := gin.Default()

		r.POST("/video", api.AuthMiddleware(), api.AddVideo)
		r.POST("/login", api.Login())
		r.POST("/register", api.Register)
		r.POST("/user/:id/playlist", api.AddPlaylist)
		r.GET("/user/:id/playlist/:playlist_id", api.GetPlaylist)
		r.GET("/user/:id/playlist", api.GetUserPlaylist)
		r.GET("/video/:id", api.GetVideo)
		r.GET("/user/:id/video", api.GetUserVideo)

		log.Printf("Starting webserver")
		r.Run()
	*/
}
