package controller

import (
	"github.com/gin-gonic/gin"
	"gorm/model"
)

func Create(g *gin.Context) {
	t := new(model.Test)
	g.JSONP(200, gin.H{"message": t.Create()})
}

func Del(g *gin.Context) {

}

func Add(g *gin.Context) {

}

func Select(g *gin.Context) {

}

func Update(g *gin.Context) {

}
