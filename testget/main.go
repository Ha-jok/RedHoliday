package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	engine := gin.Default()

	engine.GET("/hello/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		c.JSON(http.StatusOK,gin.H{
			"msg" : uid,
		})
	})

	engine.Run()
}

