package auth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Application credentials (not for the users), used for YouTube API
type oAuthCredential struct {
	Cid     string `json:"client_id"`
	Csecret string `json:"client_secret"`
}

type UserCredential struct {
	UserID     string `json: "user_id"`
	OAuthToken string `json: "oauth_token"`
}

var oautConfig *oauth2.Config
var aesKey string

func init() {
	aesKey = os.Getenv("AES_KEY")
	if aesKey == "" {
		log.Fatalf("No AES key")
	}

	file, err := ioutil.ReadFile("creds.json")
	if err != nil {
		log.Fatalf("Error reading oAuth2 credentials %s", err.Error())
	}

	var credentials oAuthCredential
	json.Unmarshal(file, &credentials)
	log.Printf("CLIENT ID: %s, CSECRET: %s", credentials.Cid, credentials.Csecret)
	oautConfig = &oauth2.Config{
		ClientID:     credentials.Cid,
		ClientSecret: credentials.Csecret,
		RedirectURL:  "http://127.0.0.1:8000/oauth",
		Scopes: []string{
			"https://www.googleapis.com/auth/youtube",
			"https://www.googleapis.com/auth/youtubepartner",
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

// Retun the URL for authentica with google OAuth
func GetOAuthURL(state string) string {
	return oautConfig.AuthCodeURL(state)
}

// Get google's OAuth token
func GetOAuthToken(code string) (*oauth2.Token, error) {
	token, err := oautConfig.Exchange(context.TODO(), code)
	if err != nil {
		return &oauth2.Token{}, err
	}
	return token, nil
}
