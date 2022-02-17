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

// Pass an empty string to use the default path "img/thumbnail"
func NewThumbnailLocalStore(basePath string) ThumbnailLocalStore {
	if basePath == "" {
		basePath = "img/thumbnail"
	}
	return ThumbnailLocalStore{basePath}
}

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
