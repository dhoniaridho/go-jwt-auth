package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid      string `gorm:"unique;not_null"`
	Name     string `gorm:"not_null"`
	Email    string `gorm:"not_null"`
	Password string `json:"-" gorm:"not_null"`
}
