package todos

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	controller := Controller{}

	r.GET("/", controller.Index)
	r.GET("/:id", controller.Show)

	return r
}
