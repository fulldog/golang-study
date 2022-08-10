package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm/model"
	"io/ioutil"
	"strconv"
	"sync"
)

func Create(g *gin.Context) {
	t := new(model.Test)
	g.JSONP(200, gin.H{"message": t.Create()})
}

func Del(g *gin.Context) {

}

func Add(g *gin.Context) {
	by, _ := ioutil.ReadAll(g.Request.Body)
	if len(by) == 0 {
		g.JSON(200, gin.H{
			"message": "没有数据",
		})
		return
	}

	sy := sync.WaitGroup{}
	sy.Add(1000)
	sl := make([]*model.Test, 1000)
	for i := 0; i < 1000; i++ {
		go func(i int) {
			t := new(model.Test)
			json.Unmarshal(by, t)
			t.Name = t.Name + strconv.Itoa(i)
			sl[i] = t
			fmt.Println(t.Add())
			sy.Done()
		}(i)
	}
	sy.Wait()
	g.JSON(200, gin.H{
		"data": sl,
	})

}

func Select(g *gin.Context) {

}

func Update(g *gin.Context) {

}
