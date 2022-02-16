package auth

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Application credentials for YouTube API
type oAuthCredential struct {
	Cid     string `json:"client_id"`
	Csecret string `json:"client_secret"`
}

// Used as payload for JWE token
type UserCredential struct {
	ID         int32         `json:"user_id"`
	OAuthToken *oauth2.Token `json:"oauth_token"`
}

// Google profile info
type UserInfo struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Picture       string `json:"picture"`
}

var oautConfig *oauth2.Config
var aesKey string

func init() {

	// Retreiving AES key from env
	aesKey = os.Getenv("AES_KEY")
	if aesKey == "" {
		log.Fatalf("No AES key")
	}

	// Register the struct oauth2.Token for http sessions
	gob.Register(&oauth2.Token{})

	// Load application credentials
	file, err := ioutil.ReadFile("creds.json")
	if err != nil {
		log.Fatalf("Error reading oAuth2 credentials %s", err.Error())
	}
	var credentials oAuthCredential
	json.Unmarshal(file, &credentials)

	oautConfig = &oauth2.Config{
		ClientID:     credentials.Cid,
		ClientSecret: credentials.Csecret,
		RedirectURL:  "http://127.0.0.1:8000/oauth",
		Scopes: []string{
			"https://www.googleapis.com/auth/youtube",
			"https://www.googleapis.com/auth/youtubepartner",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

// Hash the string using bcrypt
func Hash(plainText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), 13)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Retun the URL for authentication with google OAuth
func GetOAuthURL(state string) string {
	return oautConfig.AuthCodeURL(state)
}

// Get Google OAuth token
func GetOAuthToken(code string) (*oauth2.Token, error) {
	token, err := oautConfig.Exchange(context.TODO(), code)
	if err != nil {
		return &oauth2.Token{}, err
	}
	return token, nil
}

func GetUserInfo(token *oauth2.Token) (UserInfo, error) {

	client := oautConfig.Client(context.TODO(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return UserInfo{}, err
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	var info UserInfo
	err = json.Unmarshal(data, &info)
	if err != nil {
		return UserInfo{}, err
	}
	return info, nil
}
