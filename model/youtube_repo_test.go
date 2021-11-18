package model

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetYTRepository(t *testing.T) {
	_, errS := GetYTRepository()
	if errS != ErrNoRepoInstance {
		t.Fail()
	}
	_db, _, err := sqlmock.New()
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

	_, err = GetYTRepository()
	if err != nil {
		t.Fail()
	}
}

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

}

func TestGetVideo(t *testing.T) {

}
