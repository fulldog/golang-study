package controller

import "github.com/gin-gonic/gin"

func Create(g *gin.Context) {
	g.JSONP(200, gin.H{"message": "ok"})
}

func Del(g *gin.Context) {

}

func Add(g *gin.Context) {

}

func Select(g *gin.Context) {

}

func Update(g *gin.Context) {

}
