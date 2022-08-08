package router

import (
	"github.com/gin-gonic/gin"
	"gorm/controller"
)

func Register(g *gin.Engine) {
	g.POST("/create", controller.Create)
	g.POST("/update", controller.Update)
	g.GET("/select", controller.Select)
	g.POST("/add", controller.Add)
	g.POST("/del", controller.Del)
	g.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
