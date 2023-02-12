package main

import (
	"api/src/config/database"
	"api/src/internal/apps/users"

	internal "api/src/internal/apps"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	r := gin.Default()

	internal.Setup(r)

	database.Init()

	database.GetDb().AutoMigrate(&users.User{})

	r.Run()

}
