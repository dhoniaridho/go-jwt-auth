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

	token, err := c.authService.SignIn(&cridential)

	if err != nil {
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
		},
	})
}

func (c *Controller) Register(ctx *gin.Context) {

}
