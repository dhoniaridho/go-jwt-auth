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

func (UserService) CreateOne(user *CreateUserDto) (*CreateUserDto, error) {

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

func (UserService) UpdateOne(id int, payload *UpdateUserDto) (*UpdateUserDto, error) {
	db := database.GetDb()

	user := User{
		ID: id,
	}

	err := db.First(&user).Error

	if user.Password != "" {
		password := []byte(user.Password)

		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

		if err != nil {
			return payload, err
		}

		user.Password = string(hashedPassword)

	}

	if err != nil {
		return nil, err

	}

	user.Name = payload.Name
	user.Email = payload.Email

	savingErr := db.Save(&user).Error

	if savingErr != nil {
		return payload, savingErr
	}

	return payload, nil

}

func (UserService) DeleteOne(id int) (User, error) {
	db := database.GetDb()
	user := User{
		ID: id,
	}

	deleteError := db.Delete(&user).Error

	if deleteError != nil {
		return user, deleteError
	}

	return user, nil

}
