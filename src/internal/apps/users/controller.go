package users

import (
	"api/src/database"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserService
}

func (s *Controller) Index(ctx *gin.Context) {

	data := s.UserService.GetAll()

	ctx.JSON(
		200,
		gin.H{
			"success": true,
			"data":    data,
		},
	)

}

func (s *Controller) Show(ctx *gin.Context) {

	id := ctx.Param("id")

	data, err := s.UserService.GetOne(id)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
			"status":  404,
		})
		return
	}

	ctx.JSON(
		200,
		gin.H{
			"success": true,
			"data":    data,
		},
	)

}

func (s *Controller) Create(ctx *gin.Context) {

	var user UserDto

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": strings.Split(err.Error(), "\n"),
			"status":  400,
		})
		return
	}

	s.UserService.CreateOne(&UserDto{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	newUser := struct {
		Name  string
		Email string
	}{
		Name:  user.Name,
		Email: user.Email,
	}

	ctx.JSON(201, gin.H{
		"message": "Created",
		"data":    newUser,
	})

}

func (c *Controller) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	db := database.GetDb()

	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.String(400, "Invalid ID")
		return
	}

	user := User{
		ID: id,
	}

	deleteError := db.Delete(&user).Error

	if deleteError != nil {
		ctx.JSON(400, gin.H{
			"message": deleteError,
		})
	}

	ctx.JSON(200, gin.H{
		"message": "Successfully deleted",
		"data":    nil,
	})
}
