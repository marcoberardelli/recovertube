package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var userRepo UserRepository

func initUserRepository(dsn string) {
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error initializing yt repo")
	}
	userRepo = UserRepository{db: _db}
	//userRepo.db.AutoMigrate(&User{}, &Video{}, &Playlist{})
}

func GetUserRepository() UserRepository {
	return userRepo
}

func (r UserRepository) GetUser(id int) (User, error) {
	var user User
	res := r.db.First(&user, id)
	if res.Error != nil {
		return User{}, res.Error
	} else if res.RowsAffected == 0 {
		return User{}, ErrUserNotRegistered
	}

	return user, nil
}

func (r UserRepository) GetUserID(email string) (int32, error) {
	var user User
	res := r.db.Where("email = ?", email).Select("id").First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		return -1, ErrUserNotRegistered
	} else if res.Error != nil {
		return -1, res.Error
	}

	return user.ID, nil
}

func (r UserRepository) CreateUser(email string) (int32, error) {
	user := User{Email: email}
	if res := r.db.Select("email").Create(&user); res.Error != nil {
		return -1, res.Error
	}
	return user.ID, nil
}

func (r UserRepository) GetUserByEmail(email string) (User, error) {
	var user User
	res := r.db.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return User{}, res.Error
	} else if res.RowsAffected == 0 {
		return User{}, ErrUserNotRegistered
	}

	return user, nil
}
