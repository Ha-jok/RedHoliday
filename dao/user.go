package dao

import (
	"RedHoliday/model"
)

//创建一个新用户的数据表
func CreateNewPerson(username string) bool{
	//连接数据库
	db := Link_mysql()
	//创建数据表
	db.Create(&model.Person_mysql{
		Username: username,
	})
	return true
}
//在新用户表中插入相关信息
func InsertNewTable(username string, uid int) bool{
	db := Link_mysql()
	db.Save(&model.Person_mysql{
		Uid: uid,
		Username: username,
		Avatar: "",
		Friends: "",
	})
	return true
}

//在用户组表中插入新的用户
func InsertNewPerson(username,password,email,salt string,phone int)bool{
	//连接数据库
	db := Link_mysql()
	//插入数据
	db.Save(&model.Users_mysql{
		Username: username,
		Password: password,
		Email: email,
		Salt: salt,
	})
	return true
}


//通过用户名查询密码，
func QueryUp(username string)(pw,salt string){
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var u model.Users_mysql
	db.Where("username = ? ",username).Find(&u)
	return u.Password,u.Salt
}

//通过用户名查询uid
func QueryUid(username string)int{
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var u model.Users_mysql
	db.Where("username = ?",username).Find(&u)
	return u.Uid
}



//通过username查询个人信息表
func QueryUsernameIntroduction(username string) model.Person_mysql {
	//连接数据库
	db := Link_mysql()
	//定义参数
	var user model.Person_mysql
	db.Table(username).First(&user)
	return user
}


//通过邮箱查询用户名
func QueryEmail(email string)string{
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var u model.Users_mysql
	db.Table("username").Where("email = ?",email).Find(&u)
	return u.Username
}



//更改购物车信息
func UpdateCart(username,cart string){
	//连接数据库
	db := Link_mysql()
	//定义修改语句
	db.Table(username).Update("shopping_cart",cart)
}

//更改支付订单信息
func UpdateOrderPaid(username,paid string){
	//连接数据库
	db := Link_mysql()
	//定义修改语句
	db.Table(username).Update("order_paid",paid)
}

//更改待支付订单信息
func UpdateOrderUnpaid(username,order_un string){
	//连接数据库
	db := Link_mysql()
	db.Table(username).Update("order_unpaid",order_un)
}

//更改已收货订单信息
func UpdateOrderReceived(username,received string){
	//连接数据库
	db := Link_mysql()
	db.Table(username).Update("order_received",received)

}
