package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Name     string
	Email    string
	Password string
}
