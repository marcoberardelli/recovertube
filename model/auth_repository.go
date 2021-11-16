package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

var authRepo AuthRepository

func InitAuthRepository(dsn string) (AuthRepository, error) {

	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return AuthRepository{}, err
	}
	authRepo = AuthRepository{db: _db}
	return authRepo, nil
}

func GetAuthRepository() (AuthRepository, error) {
	if authRepo.db == nil {
		return AuthRepository{}, ErrMultipeRepoInstance
	}
	return authRepo, nil
}

func (r AuthRepository) Register(email, password string) error {

	return nil
}
func (r AuthRepository) Login(email, password string) (string, error) {

	return "", nil
}
func (r AuthRepository) Logout(user User) error {

	return nil
}
