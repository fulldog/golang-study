package main

import (
	"github.com/gin-gonic/gin"
	"gorm/router"
)

func main() {
	g := gin.New()

	g.Use(gin.Logger())

	g.Use(gin.Recovery())

	router.Register(g)

	g.Run(":8080")
}
