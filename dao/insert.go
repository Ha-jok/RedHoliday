package dao

import "fmt"

//该go文件存放插入数据库函数

func Insert_new_person(username,password,email,salt string,phone int)bool{
	//定义参数
	db := Link_mysql()
	insertstring := "insert into users(username,password,email,salt,phone) values(?,?,?,?,?)"
	//插入数据
	_,err := db.Exec(insertstring,username,password,email,salt,phone)
	//返回信息
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}


func Insert_new_table(username string, uid int) bool{
	db := Link_mysql()
	insertstring := "insert into "+username+"(username,uid,avatar,friends,follow_business,shopping_cart,order_paid,order_unpaid,order_received) values(?,?,?,?,?,?,?,?,?)"
	//插入数据
	_,err := db.Exec(insertstring,username,uid,"无","无","无","无","无","无","无")
	//返回信息
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
