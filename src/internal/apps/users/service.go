package users

import (
	"api/src/config/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/jaevor/go-nanoid"
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

	err := db.Where("Uid = ?", id).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return User{}, err
	}

	return user, err
}

func (UserService) FindOneByEmail(email string) (User, error) {

	db := database.GetDb()

	var user User

	err := db.Where("email = ?", email).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return User{}, err
	}

	return user, err
}

func (UserService) CreateOne(user *CreateUserDto) (*User, error) {

	db := database.GetDb()

	password := []byte(user.Password)
	id, _ := nanoid.Standard(21)

	validateEmail := func(email string) bool {
		var count int64
		db.Model(&User{}).Where("email = ?", email).Count(&count)
		return count == 0
	}

	// Validate the email address
	if !validateEmail(user.Email) {
		return &User{}, errors.New("email is not available")
	}

	idUser := id()

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return &User{
			Uid:   idUser,
			Name:  user.Name,
			Email: user.Email,
		}, err
	}

	u := &User{
		Uid:      idUser,
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	db.Create(u)

	return u, nil
}

func (UserService) UpdateOne(id string, payload *UpdateUserDto) (*UpdateUserDto, error) {
	db := database.GetDb()

	user := User{
		Uid: id,
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

func (UserService) DeleteOne(id string) (User, error) {
	db := database.GetDb()
	user := User{
		Uid: id,
	}

	deleteError := db.Delete(&user).Error

	if deleteError != nil {
		return user, deleteError
	}

	return user, nil

}
