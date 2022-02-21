package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main(){
	//设置连接语句
	sqlstr := "root:2002@tcp(127.0.0.1:3306)/test"
	//初始化连接
	db,err := sql.Open("mysql",sqlstr)
	if err != nil {
		log.Fatal(err.Error())
	}
	//检查连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	//创建数据表
	//in := "test"
	//_,err = db.Exec("create table `"+in+"` (`id` int(20) auto_increment,`username` varchar(20),primary key (`id`))")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	//插入数据
	insertstring := "insert into user(username) values(?)"
	_,err = db.Exec(insertstring,"test1")


	//查询数据
	type user struct {
		id int
		username string
	}
	var u user
	querystring := "select * from user where username = ?"
	rows,err := db.Query(querystring,"2")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next(){
		err = rows.Scan(&u.id,&u.username)
		if err != nil {
			return
		}

	}
	fmt.Println(u.username)

	//替换数据
	updatestr := "update user set username=? where id = ?"
	_, err = db.Exec(updatestr, "testup", 1)
	if err != nil {
		log.Fatal(err.Error())
	}









}