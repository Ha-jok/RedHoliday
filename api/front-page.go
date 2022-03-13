package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//该go文件包含了front-page下接口：主页


func FrontPage(c *gin.Context){
	//验证token
	username := c.MustGet("username").(string)
	//token有效，则提取token中的用户名
	if username != "" {
		c.JSON(http.StatusOK,gin.H{
			"message" : "欢迎回来"+username,
		})
		return
	}
	//token无效,
	c.JSON(http.StatusOK,gin.H{
		"message" : "欢迎回来游客",
	})
}