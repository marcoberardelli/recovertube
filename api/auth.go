package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"
)

const AES_KEY = "AES_KEY_TEST"

var encrypter jose.Encrypter

type LoginForm struct {
	email    string `form:"email"`
	password string `form:"password"`
}

type RegisterForm struct {
	email           string `form:"email"`
	password        string `form:"password"`
	confirmPassword string `form"confirm_password`
}

func init() {
	var err error
	encrypter, err = jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.A256KW, Key: AES_KEY}, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(405, "No authorization token provided")
		}

		parsedToken, err := jose.ParseEncrypted(token)
		if err != nil {
			log.Printf("Error parsing token [%s] : %s", token, err)
			c.JSON(500, "Error parsing auth token")
		}

		decryptedToken, err := parsedToken.Decrypt(AES_KEY)
		if err != nil {
			log.Printf("Error decrypting token [%s] : %s", token, err)
			c.JSON(500, "Error decrypting token")
		}

		//TODO: define decryptedToken
		c.Set("User", decryptedToken)
		c.Next()
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		credentials := LoginForm{}
		err := c.ShouldBind(&credentials)
		if err != nil {
			log.Printf("Error parsing login form %s", err)
			c.JSON(400, "Form wrong")
		}

		//TODO: salt password
		//model.CompareLogin(credentials.email, credentials.password)

		json := fmt.Sprintf(`{user: "%s"}`, credentials.email)
		obj, err := encrypter.Encrypt([]byte(json))
		if err != nil {
			log.Printf("Error encrypting credentials %s", err)
			c.JSON(500, "Error encrypting credentials")
		}
		serialized := obj.FullSerialize()

		c.JSON(200, fmt.Sprintf(`{Token:"%s"}`, serialized))
	}

}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		registerForm := RegisterForm{}
		err := c.ShouldBind(&registerForm)
		if err != nil {
			log.Printf("Error processing the form %s", err)
			c.JSON(500, "Error processing the form ")
		}

	}
}
