package database

import (
	env "api/src/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {

	dsn := env.Get("DB_MYSQL_CONNECTION")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connection to database failed")
	}

	db = database

	return err

}

func GetDb() *gorm.DB {
	return db
}
