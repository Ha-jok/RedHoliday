package dao

import (
	"RedHoliday/model"
	"fmt"
	"time"
)

//该go文件下存放查询数据库的函数

//通过用户名查询密码，
func Query_up(username string)(pw,salt string){
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
func Query_uid(username string)int{
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
func Query_username_introduction(username string) model.Person_mysql {
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
func Query_email(email string)string{
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var u model.Users_mysql
	querystr := "select username from users where email=?"
	//查询信息
	_ = db.QueryRow(querystr,email).Scan(&u.Username)
	return u.Username
}


//通过uid查询商品详细信息
func Query_commmidty(uid int)model.Commidity{
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var c model.Commidity
	querystr := "select * from commidity where uid=?"
	//查询信息
	r := db.QueryRow(querystr,uid).Scan(&c.Uid,&c.Commidity_name,&c.Volume,&c.Evaluations,&c.Detailed_Introduction)
	fmt.Println(r)
	return c
}


//查询所有商品信息
func Query_commiditys()map[int]string{
	var commiditys = make(map[int]string,20)
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var commidity model.Commidity
	querystr := "select uid,name from commidity where uid >0"
	//查询信息
	rows, err := db.Query(querystr)
	if err != nil {
		fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
		return commiditys
	}
	for rows.Next() {
		err := rows.Scan(&commidity.Uid,&commidity.Commidity_name)
		if err != nil {
			fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
			return commiditys
		}
		commiditys[commidity.Uid] = commidity.Commidity_name
	}
	return commiditys
}