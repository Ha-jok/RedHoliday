package main

import (
	"RedHoliday/dao"
	"RedHoliday/model"
	"RedHoliday/service"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

func main(){
	engine := gin.Default()
	engine.POST("/regist", func(c *gin.Context) {



			//绑定参数并处理错误
			var New_Person model.New_person
			err := c.ShouldBind(&New_Person)
			if err != nil{
				log.Fatal(err.Error())
			}
			//检验参数是否合法
			//验证用户名
			if len(New_Person.Username)>6||len(New_Person.Username)==0{
				c.JSON(http.StatusOK,gin.H{
					"message" : "用户名格式错误",
				})
				return
			}
			//验证密码
			if len(New_Person.Password)<6||len(New_Person.Password)>12{
				c.JSON(http.StatusOK,gin.H{
					"message" : "密码格式错误",
				})
				return
			}
			//验证电话号码
			if len(strconv.Itoa(New_Person.Phone)) != 11{
				c.JSON(http.StatusOK,gin.H{
					"message" : "电话号码格式错误",
				})
				return
			}
			//验证邮箱
		    b := service.VerifyEmailFormat(New_Person.Email)
			if !b {
				c.JSON(http.StatusOK,gin.H{
					"message" : "邮箱格式不对",
				})
				return
			}
			//发送验证码



			//检验验证码是否正确



			//储存到数据库,添加到用户总表，并创建一个数据表
			//储存到数据库
			b = dao.Insert_new_person(New_Person.Username,New_Person.Password,New_Person.Email,New_Person.Phone)
			if !b {
				c.JSON(http.StatusOK,gin.H{
					"message" : "注册失败",
				})
			}
			//为用户创建一个个人数据表
			b = dao.Create_new_person(New_Person.Username)

			//返回数据
			if b {
				c.JSON(http.StatusOK,gin.H{
					"message" : "注册成功"+New_Person.Username,
				})
				return
			}
			c.JSON(http.StatusOK,gin.H{
				"message" : "注册失败",
			})






	})
	engine.Run()
}