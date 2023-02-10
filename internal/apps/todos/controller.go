package todos

import "github.com/gin-gonic/gin"

type Controller struct {
	Service
}

func (this *Controller) Index(ctx *gin.Context) {

	data := this.Service.GetAll()

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    data,
	})
}

func (this *Controller) Show(ctx *gin.Context) {
	id := ctx.Param("id")

	data := this.Service.GetOne(id)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    data,
	})
}
