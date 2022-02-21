package dao

import "fmt"

//该go文件存放创建数据表操作的函数

func Create_new_person(username string) bool{
	//连接数据库
	db := Link_mysql()
	//创建数据表
	_,err := db.Exec("create table`"+username+"`(`uid` int(10),`username` varchar(30),`avatar` varchar(30),`friends` varchar(30),`follow_business` text,`shopping_cart` text,`order_paid` text,`order_unpaid` text,`order_received` text,primary key (`uid`))charset=utf8")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

