package main

import (
	"RedHoliday/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//两个接口，一个发送验证码一个验证

func main(){
	engine := gin.Default()
	var verify_code_send string
	engine.POST("/email/verify", func(c *gin.Context) {
		email := c.PostForm("email")
		//判断邮箱格式
		b := service.VerifyEmailFormat(email)
		if !b {
			c.JSON(http.StatusOK,gin.H{
				"message" : "邮箱格式错误",
			})
			return
		}
		verify_code_send = service.Email_verify_code(email)
	})
	engine.POST("/email",func(c *gin.Context) {
		verify_code := c.PostForm("verify_code")
		//获取上下文中的验证码
		if verify_code_send == verify_code{
			c.JSON(http.StatusOK,gin.H{
				"message" : "c",
			})
			return
		}
		//验证码错误
	})
	engine.Run()
}


