package main

import (
	"log"
	"recovertube/route"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine) {

	r.GET("/login", route.OAuthLogin())
	r.GET("/oauth", route.OAuthCallback())

	r.GET("/video", route.AuthMiddleware(), route.GetAllVideos)
	r.GET("/video/:id", route.AuthMiddleware(), route.GetVideo)
	r.POST("/playlist/:playlist_id/video/video_id", route.AuthMiddleware(), route.AddVideo)

	r.GET("playlist", route.AuthMiddleware(), route.GetAllPlaylists)
	r.GET("playlist/:playlist_id", route.AuthMiddleware(), route.GetPlaylist)
	r.POST("/playlist", route.AuthMiddleware(), route.NewPlaylist)
	r.POST("/playlist/:playlist_id", route.AuthMiddleware(), route.AddPlaylist)
}

func main() {

	log.SetFlags(log.LstdFlags)

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Domain: "127.0.0.1"})
	r.Use(sessions.Sessions("session", store))

	registerRoutes(r)

	log.Printf("Starting webserver")
	r.Run()

}
