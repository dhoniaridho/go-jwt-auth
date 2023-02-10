package users

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	c := Controller{}

	r.GET("/", c.Index)
	r.GET("/:id", c.Show)

	return r
}
