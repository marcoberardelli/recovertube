package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"recovertube/api"

	"github.com/gin-gonic/gin"
)

func main() {

	log.SetFlags(log.LstdFlags)

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Domain: "127.0.0.1"})

	r.Use(sessions.Sessions("mysession", store))

	r.POST("/video", api.AuthMiddleware(), api.AddVideo)
	r.GET("/login", api.OAuthLogin())
	r.GET("/oauth", api.OAuthCallback())
	r.POST("/user/:id/playlist", api.AddPlaylist)
	r.GET("/user/:id/playlist/:playlist_id", api.GetPlaylist)
	r.GET("/user/:id/playlist", api.GetUserPlaylist)
	r.GET("/video/:id", api.GetVideo)
	r.GET("/user/:id/video", api.GetUserVideo)

	log.Printf("Starting webserver")
	r.Run()

}
