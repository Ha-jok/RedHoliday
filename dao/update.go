package dao

import (
	"fmt"
	"time"
)

//该go文件存放更改信息操作

//更改购物车信息
func Update_cart(username,cart string){
	//连接数据库
	db := Link_mysql()
	//定义修改语句
	updatestr := "update "+username+" set shopping_cart=? where username = ?"
	//修改数据库并处理错误
	_, err := db.Exec(updatestr, cart,username)
	if err != nil {
		fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
		return
	}
}

//更改支付订单信息
func Update_Order_paid(username,paid string){
	//连接数据库
	db := Link_mysql()
	//定义修改语句
	updatestr := "update "+username+" set order_paid=? where username = ?"
	//修改数据库并处理错误
	_, err := db.Exec(updatestr, paid,username)
	if err != nil {
		fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
		return
	}
}

//更改待支付订单信息
func Update_order_unpaid(username,order_un string){
	//连接数据库
	db := Link_mysql()
	//定义修改语句
	updatestr := "update "+username+" set order_unpaid=? where username = ?"
	//修改数据库并处理错误
	_, err := db.Exec(updatestr, order_un,username)
	if err != nil {
		fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
		return
	}
}

//更改已收货订单信息
func Update_order_received(username,received string){
	//连接数据库
	db := Link_mysql()
	//定义修改语句
	updatestr := "update "+username+" set order_received=? where username = ?"
	//修改数据库并处理错误
	_, err := db.Exec(updatestr, received,username)
	if err != nil {
		fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
		return
	}
}

//评论商品
func Update_ecalutions(uid int,comment string){
	//连接数据库
	db := Link_mysql()

	updatestr := "update commidity set evaluations=? where uid = ?"
	_, err := db.Exec(updatestr, comment,uid)
	if err != nil {
		return
	}
}
