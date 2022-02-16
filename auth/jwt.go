package auth

import (
	"bytes"
	"encoding/gob"
	"log"

	"gopkg.in/square/go-jose.v2"
)

func EncryptToken(userID, token string) (string, error) {
	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.A256KW, Key: aesKey}, nil)
	if err != nil {
		log.Printf("Error creating jose Encrypter")
		return "", err
	}

	payload := UserCredential{
		UserID:     userID,
		OAuthToken: token,
	}
	// Converting JwePayload struct to []byte
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err = enc.Encode(payload)
	if err != nil {
		log.Printf("Error creating the bytes encoder")
		return "", err
	}
	jwe, err := encrypter.Encrypt(buf.Bytes())
	if err != nil {
		log.Printf("Error encryping bytes")
		return "", err
	}

	return jwe.CompactSerialize()
}

func DecryptToken(token string) (UserCredential, error) {
	parsedToken, err := jose.ParseEncrypted(token)
	if err != nil {
		log.Printf("Error parsing token [%s] : %s", token, err)
		return UserCredential{}, err
	}

	decryptedToken, err := parsedToken.Decrypt(aesKey)
	if err != nil {
		log.Printf("Error decrypting token [%s] : %s", token, err)
		return UserCredential{}, err
	}

	//TODO: check if correct or random string
	buf := bytes.NewReader(decryptedToken)
	payload := UserCredential{}
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&payload)
	if err != nil {
		return UserCredential{}, err
	}

	return payload, err
}
