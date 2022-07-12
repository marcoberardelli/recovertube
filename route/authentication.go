// Copyright (C) 2022  Marco Berardelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package route

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

func getUser(c *gin.Context) (auth.UserCredential, bool) {
	userInterface, ok := c.Get("User")
	if !ok {
		return auth.UserCredential{}, false
	}
	user, ok := userInterface.(auth.UserCredential)
	return user, ok
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

func clearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// Start login
func OAuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Preparing state and store in session
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

// Callback called after logged in with Google API
func OAuthCallback() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// Compare state from session and the one passed as http call parameter
		stateQuery := c.Query("state")
		stateSess := session.Get("state")
		if stateQuery != stateSess {
			clearSession(c)
			c.JSON(500, "Wrong state")
			return
		}

		// Generating Google OAuth token
		code := c.Query("code")
		token, err := auth.GetOAuthToken(code)
		if err != nil {
			log.Printf("Error retreiving the token in OAuthCallback\n")
			clearSession(c)
			c.JSON(500, "Error retreiving Google OAuth token")
			return
		}
		info, err := auth.GetUserInfo(token)
		if err != nil {
			clearSession(c)
			c.JSON(500, "Error querying for user info")
			return
		}
		if !info.EmailVerified {
			clearSession(c)
			c.JSON(409, "Please confirm your Google account before continuing")
			return
		}
		_id, err := model.GetUserID(info.Email)
		id := &_id
		if err == model.ErrUserNotRegistered {
			// Creates a new user
			_id, err = model.CreateUser(info.Email)
			if err != nil {
				clearSession(c)
				c.JSON(500, "Error creating user")
				return
			}
		} else if err != nil {
			clearSession(c)
			log.Printf("Error %s", err.Error())
			c.JSON(500, "Internal erro")
			return
		}

		// Generating JWE token
		jwe, err := auth.EncryptJWE(auth.UserCredential{ID: *id, OAuthToken: token})
		if err != nil {
			clearSession(c)
			c.JSON(500, "Internal error E")
			return
		}
		c.JSON(200, gin.H{"token": jwe})
	}

}
