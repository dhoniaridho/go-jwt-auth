package users

import "github.com/gin-gonic/gin"

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
