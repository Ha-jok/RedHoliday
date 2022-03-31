package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

//该go文件存放链接数据库函数
func Link_mysql() *gorm.DB {
	//设置连接语句
	sqlstr := "root:2002@tcp(127.0.0.1:3306)/redholiday"
	//初始化连接
	db,err := gorm.Open(mysql.Open(sqlstr),&gorm.Config{})
	if err != nil{
		panic(err)
		return db
	}
	//检查连接
	return db
}
