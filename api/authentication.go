package api

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"log"
	"recovertube/auth"
	"recovertube/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
)

func init() {
	gob.Register(&oauth2.Token{})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenJWE := c.Request.Header.Get("Authorization")
		if tokenJWE == "" {
			c.JSON(405, "No authorization token provided")
		}
		token, err := auth.DecryptJWE(tokenJWE)
		if err != nil {
			log.Printf("Error decrypting JWE\n")
			c.JSON(500, "Internal error")
		}
		c.Set("User", token)
		c.Next()
	}
}

func OAuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		b := make([]byte, 32)
		rand.Read(b)
		state := base64.StdEncoding.EncodeToString(b)
		session := sessions.Default(c)
		session.Set("state", state)
		err := session.Save()
		if err != nil {
			log.Printf("Error saving session in OAuthLogin\n")
			c.JSON(500, "Internal error")
		}

		link := auth.GetOAuthURL(state)
		log.Printf("Link OAuth: %s\n", link)
		c.JSON(200, gin.H{"link": link})
	}
}

func OAuthCallback() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		stateQuery := c.Query("state")
		stateSess := session.Get("state")
		if stateQuery != stateSess {
			session.Clear()
			session.Save()
			c.JSON(500, "Wrong state")
			return
		}

		code := c.Query("code")
		token, err := auth.GetOAuthToken(code)
		if err != nil {
			log.Printf("Error retreiving the token in OAuthCallback\n")
			c.JSON(500, "Error retreiving Google OAuth token")
			return
		}
		info, err := auth.GetUserInfo(token)
		if err != nil {
			c.JSON(500, "Error querying for user info")
			return
		}
		if !info.EmailVerified {
			session.Clear()
			session.Save()
			c.JSON(409, "Please confirm your Google account before continuing")
			return
		}
		userRepo := model.GetUserRepository()
		_id, err := userRepo.GetUserID(info.Email)
		id := &_id
		if err == model.ErrUserNotRegistered {
			_id, err = userRepo.CreateUser(info.Email)
			if err != nil {
				c.JSON(500, "Error creating user")
				return
			}
		} else if err != nil {
			log.Printf("Error %s", err.Error())
			c.JSON(500, "Internal erroADSAr")
			return
		}

		jwe, err := auth.EncryptJWE(auth.UserCredential{ID: *id, OAuthToken: token})
		if err != nil {
			log.Printf("%s", err.Error())
			c.JSON(500, "Internal error E")
			return
		}
		c.JSON(200, gin.H{"token": jwe})
	}

}
