package main

import (
	"api/src/config/database"
	internal "api/src/internal/apps"
	"api/src/internal/apps/users"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	internal.Setup(r)

	database.Init()

	r.Run()

	database.GetDb().AutoMigrate(&users.User{})

}
