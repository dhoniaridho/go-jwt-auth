package users

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService
}

func (s *Controller) Index(ctx *gin.Context) {

	data := s.userService.GetAll()

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

	data, err := s.userService.GetOne(id)

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

	var user CreateUserDto

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": strings.Split(err.Error(), "\n"),
			"status":  400,
		})
		return
	}

	s.userService.CreateOne(&CreateUserDto{
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

	c.userService.DeleteOne(idStr)

	ctx.JSON(200, gin.H{
		"message": "Successfully deleted",
		"data":    nil,
	})
}

func (c *Controller) Update(ctx *gin.Context) {

	var user UpdateUserDto

	idStr := ctx.Param("id")

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": strings.Split(err.Error(), "\n"),
			"status":  400,
		})
		return
	}

	updated, updateErr := c.userService.UpdateOne(idStr, &UpdateUserDto{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if updateErr != nil {
		ctx.JSON(400, gin.H{
			"message": updateErr,
			"status":  400,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"status":  200,
		"data":    updated,
	})

}
