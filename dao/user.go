package dao

import (
	"RedHoliday/model"
	"fmt"
	"time"
)

//创建一个新用户的数据表
func CreateNewPerson(username string) bool{
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
//在新用户表中插入相关信息
func InsertNewTable(username string, uid int) bool{
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

//在用户组表中插入新的用户
func InsertNewPerson(username,password,email,salt string,phone int)bool{
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


//通过用户名查询密码，
func QueryUp(username string)(pw,salt string){
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var u model.Users_mysql
	querystr := "select password,salt from users where username=?"
	//查询信息
	_ = db.QueryRow(querystr,username).Scan(&u.Password,&u.Salt)
	return u.Password,u.Salt
}

//通过用户名查询uid
func QueryUid(username string)int{
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var u model.Users_mysql
	querystr := "select uid from users where username=?"
	//查询信息
	_ = db.QueryRow(querystr,username).Scan(&u.Uid)
	return u.Uid
}



//通过username查询个人信息表
func QueryUsernameIntroduction(username string) model.Person_mysql {
	//连接数据库
	db := Link_mysql()
	//定义参数
	var user model.Person_mysql
	querystring := "select * from "+username+";"
	//提取信息
	r := db.QueryRow(querystring).Scan(&user.Uid, &user.Username,&user.Avatar,&user.Friends,&user.Follow_business,  &user.Shopping_cart,&user.Order_paid,&user.Order_unpaid,&user.Order_received)
	fmt.Println(r,user)
	return user
}


//通过邮箱查询用户名
func QueryEmail(email string)string{
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var u model.Users_mysql
	querystr := "select username from users where email=?"
	//查询信息
	_ = db.QueryRow(querystr,email).Scan(&u.Username)
	return u.Username
}



//更改购物车信息
func UpdateCart(username,cart string){
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
func UpdateOrderPaid(username,paid string){
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
func UpdateOrderUnpaid(username,order_un string){
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
func UpdateOrderReceived(username,received string){
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
