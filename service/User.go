package service

import (
	"RedHoliday/dao"
	"RedHoliday/model"
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"time"
)

//该go文件存放/user路径下所调用的非jwt服务函数



//判断密码是否正确
func Judge_up(un,pw string)bool{
	//提取盐值和加密后的账号
	upw,salt := dao.Query_up(un)
	//对用户输入密码进行加密
	pw1 := pw+salt
	pwj := Encry_pw(pw1)
	var b bool
	if upw == pwj {
		b = true
	}
	return b
}

//判断邮箱格式是否正确
//代码参考CSDN，https://blog.csdn.net/daimading/article/details/88390302
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}



//创建新用户
func Create_user(username,password,email,salt string,phone int)bool{
	//将参数传入数据库,储存到用户表
	b := dao.Insert_new_person(username,password,email,salt,phone)
	if !b {
		return false
	}
	//获取用户名
	uid := dao.Query_uid(username)
	//为用户创建一个个人数据表
	b = dao.Create_new_person(username)
	if !b {
		return false
	}
	//将信息储存到新表中
	b = dao.Insert_new_table(username,uid)
	if !b {
		return false
	}
	return true

}

//根据username查询个人信息（购物车，订单不可查询）
func Query_user_intruduction(username string)model.Person_mysql{


	//从数据库中提取信息
	user := dao.Query_username_introduction(username)


	//将信息返回
	return user


}


//注册用户时获取盐值
func Create_salt()string{
	//获取当前时间作为盐值
	salt := time.Now().Format("15:04:05")

	//返回时间戳
	return salt

}


//使用md5加密用户密码
func Encry_pw(pw string)string{
	h := md5.New()
	h.Write([]byte(pw))
	return hex.EncodeToString(h.Sum(nil))
}


//修改用户购物车商品,重新储存购物车信息
func Shopping_cart_revise(username,shopping_carts string){
	//重新储存用户购物车
	dao.Update_cart(username,shopping_carts)
}

//储存支付订单
func Order_paid(username,settlement string){
	//提取原有支付订单
	user := Query_user_intruduction(username)
	paid := user.Order_paid+settlement+","
	//更改用户订单状态
	dao.Update_Order_paid(username,paid)
}

//修改待支付订单
func Order_unpaid(username,order_un string){
	//更改用户订单状态
	dao.Update_order_unpaid(username,order_un)
}

//修改确认收货订单
func Order_received(username,receit string){
	dao.Update_order_received(username,receit)
}