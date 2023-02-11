package users

import (
	"api/src/database"

	"gorm.io/gorm"
)

type UserService struct{}

func (s UserService) GetAll() []User {

	db := database.GetDb()

	var users []User

	db.Find(&users)

	return users
}

func (UserService) GetOne(id string) (User, error) {

	// user := User{Name: "dhoniaridho", ID: 1}

	db := database.GetDb()

	var user User

	err := db.Where("id = ?", id).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return User{}, err
	}

	return user, err
}
