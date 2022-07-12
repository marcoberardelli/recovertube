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
	"gorm.io/gorm"
)

func GetUser(id int) (User, error) {
	var user User
	res := db.First(&user, id)
	if res.Error != nil {
		return User{}, res.Error
	} else if res.RowsAffected == 0 {
		return User{}, ErrUserNotRegistered
	}

	return user, nil
}

func GetUserID(email string) (int32, error) {
	var user User
	res := db.Where("email = ?", email).Select("id").First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		return -1, ErrUserNotRegistered
	} else if res.Error != nil {
		return -1, res.Error
	}

	return user.ID, nil
}

func CreateUser(email string) (int32, error) {
	user := User{Email: email}
	if res := db.Select("email").Create(&user); res.Error != nil {
		return -1, res.Error
	}
	return user.ID, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	res := db.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return User{}, res.Error
	} else if res.RowsAffected == 0 {
		return User{}, ErrUserNotRegistered
	}

	return user, nil
}
