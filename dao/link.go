package dao

import (
	"database/sql"
	"log"
)

//该go文件存放链接数据库函数
func Link_mysql()*sql.DB {
	//设置连接语句
	sqlstr := "root:2002@tcp(127.0.0.1:3306)/redholiday"
	//初始化连接
	db,err := sql.Open("mysql",sqlstr)
	if err != nil{

		return db
	}
	//检查连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
