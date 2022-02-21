package main

import (
	"RedHoliday/service"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"net/http"
)


func main (){
	engine := gin.Default()
	engine.GET("/:username", func(c *gin.Context) {


		//获取username
		username := c.Param("username")

		//从数据库中提取信息
		user := service.Query_user_intruduction(username)



		//返回信息
		c.JSON(http.StatusOK,gin.H{

				"uid" : user.Uid,
				"username" : user.Username,
				"friends" : user.Friends,
				"follow_business" : user.Follow_business,

		})



	})

	engine.Run()
}
