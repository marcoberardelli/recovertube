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
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var videos = []Video{Video{ID: "dQw4w9WgXcQ", Title: "Rick Astley - Never Gonna Give You Up (Official Music Video)"}}

func TestAddVideo(t *testing.T) {
	_db, mock, err := sqlmock.New()
	sqlmock.NewWithDSN(os.Getenv("PSQL_DSN"))
	if err != nil {
		t.Fail()
	}
	dialector := postgres.New(postgres.Config{
		Conn:       _db,
		DriverName: "postgres",
	})
	db, err := gorm.Open(dialector, &gorm.Config{PrepareStmt: false})
	if err != nil {
		t.Fail()
	}
	ytRepo = YoutTubeDBRepository{db: db}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO video ")

}

func TestGetVideo(t *testing.T) {

}

func TestIsVideoAvailable(t *testing.T) {

}

func TestAddPlaylist(t *testing.T) {

}

func TestGetPlaylist(t *testing.T) {

}
