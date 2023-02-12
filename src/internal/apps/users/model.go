package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid      string
	Name     string
	Email    string
	Password string `json:"-"`
}
