package users

import (
	"api/src/database"

	"golang.org/x/crypto/bcrypt"
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

	db := database.GetDb()

	var user User

	err := db.Where("id = ?", id).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return User{}, err
	}

	return user, err
}

func (UserService) CreateOne(user *UserDto) (*UserDto, error) {

	db := database.GetDb()

	password := []byte(user.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return user, err
	}

	db.Create(&User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	})

	return user, nil
}
