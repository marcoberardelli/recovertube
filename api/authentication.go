package api

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"recovertube/auth"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenJWE := c.Request.Header.Get("Authorization")
		if tokenJWE == "" {
			c.JSON(405, "No authorization token provided")
		}
		token, err := auth.DecryptToken(tokenJWE)
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
			log.Printf("stateQuery: %s\t stateSess: %s\n", stateQuery, stateSess)
			c.JSON(500, "Internal error")
		}

		code := c.Query("code")
		token, err := auth.GetOAuthToken(code)
		if err != nil {
			log.Printf("Error retreiving the token in OAuthCallback\n")
			c.JSON(500, "Internal error")
		}
		session.Set("token", token.AccessToken)
		session.Set("token-expiry", token.Expiry)
		session.Set("token-type", token.TokenType)
		session.Set("token-refresh", token.RefreshToken)
		if err = session.Save(); err != nil {
			log.Printf("%s", err.Error())
			c.JSON(500, "Error saving session")
		}
	}

}
