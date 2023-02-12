package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	authService AuthService
}

func (c *Controller) Login(ctx *gin.Context) {

	var cridential LoginDto

	if err := ctx.ShouldBind(&cridential); err != nil {
		ctx.JSON(400, gin.H{
			"message": strings.Split(err.Error(), "\n"),
			"status":  400,
		})
		return
	}

	data, err := c.authService.SignIn(&cridential)

	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
			"status":  401,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Authentication successful",
		"data":    data,
	})
}

func (c *Controller) Register(ctx *gin.Context) {
	var payload RegisterDto

	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(400, gin.H{
			"message": strings.Split(err.Error(), "\n"),
			"status":  400,
		})
		return
	}

	token, registerErr := c.authService.Register(&payload)

	if registerErr != nil {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
			"status":  401,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Authentication successful",
		"data": gin.H{
			"token": token,
			"user":  payload,
		},
	})

}
