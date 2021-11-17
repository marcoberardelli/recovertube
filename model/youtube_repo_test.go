package model

import (
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setup() (sqlmock.Sqlmock, error) {

	videos := []Video{
		{ID: "dQw4w9WgXcQ", Title: "Rick Astley - Never Gonna Give You Up (Official Music Video)", Channel: "UCuAXFkgsw1L7xaCfnd5JJOw", Available: true, LastUpdate: time.Now(), ImagePath: "images/dQw4w9WgXcQ"},
		{ID: "9bZkp7q19f0", Title: "PSY - GANGNAM STYLE(강남스타일) M/V", Channel: "UCrDkAvwZum-UTjHmzDI2iIw", Available: true, LastUpdate: time.Now(), ImagePath: "images/9bZkp7q19f0"},
	}

	err := os.Remove("test.db")
	if err != nil {
		return nil, err
	}

	_db, mock, err := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       _db,
		DriverName: "postgres",
	})
	postresDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return mock, err
	}
	ytRepo = YoutTubeDBRepository{db: postresDB}
	return mock, nil
}

func TestGetYTRepository(t *testing.T) {
	_, err := GetYTRepository()
	if err != ErrNoRepoInstance {
		t.Fail()
	}

	setup()

	_, err = GetYTRepository()
	if err != nil {
		t.Fail()
	}
}

func TestGetVideo(t *testing.T) {

}
