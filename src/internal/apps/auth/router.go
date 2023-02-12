package auth

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	c := Controller{}

	r.POST("/login", c.Login)
	r.POST("/register", c.Register)

	return r
}
