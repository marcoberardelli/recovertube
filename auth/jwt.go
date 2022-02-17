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

package auth

import (
	"bytes"
	"encoding/gob"
	"log"

	"gopkg.in/square/go-jose.v2"
)

func EncryptJWE(UserCredential UserCredential) (string, error) {
	encrypter, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.A256KW, Key: aesKey}, nil)
	if err != nil {
		log.Printf("Error creating jose Encrypter")
		return "", err
	}

	// Converting JwePayload struct to []byte
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err = enc.Encode(UserCredential)
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

func DecryptJWE(token string) (UserCredential, error) {
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
