package main

import (
	"RedHoliday/dao"
	"RedHoliday/model"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)


func main(){
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		//连接数据库
		db := dao.Link_mysql()
		//定义相关字段
		var u model.Users_mysql
		querystr := "select uid from users where password=?"
		//查询信息
		_ = db.QueryRow(querystr,"test").Scan(&u.Uid)
		fmt.Println(u.Uid)
		c.JSON(http.StatusOK,gin.H{
			"message" : u.Uid,
		})
	})
	engine.Run()


}
