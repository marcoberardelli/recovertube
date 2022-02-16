package model

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ThumbnailLocalStore struct {
	basePath string
}

var imgLocalStore = ThumbnailLocalStore{"img/thumbnail"}

func (l ThumbnailLocalStore) SaveImage(url, filename string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Printf("Get image Wrong status code")
		return errors.New("wrong status code")
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", l.basePath, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)

	return err
}

func (l ThumbnailLocalStore) GetImage(filename string) error {

	return nil
}
