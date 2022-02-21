package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//该go文件存放各种结构体

type Person_mysql struct {
	Uid int
	Username string
	Avatar string
	Friends string
	Follow_business string
	Shopping_cart string
	Order_paid string
	Order_unpaid string
	Order_received  string
}
type Users_mysql struct {
	Uid int
	Username string
	Password string
	Phone int
	Email string
	Salt string
}

type Commidity struct {
	Uid int
	Commidity_name string
	Volume int
	Evaluations string
	Detailed_Introduction string
}



type New_person struct {
	Username string `form:"Username"`
	Password string `form:"Password"`
	Phone int `form:"Phone"`
	Email string `form:"Email"`
}



type Person_login_pw struct{
	UserName string `form:"Username"`
	Password string `form:"Password"`

}



type Person_login_email struct {
	Email string `form:"Email"`
	Verify_code string `form:"Verify_code"`
}

type Person_forget_password struct {
	Phone string `form:"Phone"`
	Verify_code string `form:"Verify_code"`
	New_password string `form:"New_password"`
}


type Person_introduction struct {
	UID  int `form:"UID"`
	Username string `form:"Username"`
	Balance int  `form:"Balance"`   //余额
	Friends []string  `form:"Friends"` //[]Friend; 好友切片
	Follow_business []string  `form:"Follow_business"` //[]Follow-business;关注商家切片
}


type Person_shopping_cart struct {
	Shopping_cart []string //[]Shopping-cart;购物车切片
}


type Person_order struct {
	Order_paid []string `form:"Order_paid"`//[]Order-paid;已支付订单切片
	Order_unpaid []string `form:"Order_unpaid"` //[]Order-unpaid;未支付切片
	Pending_payment []string `form:"Pending_payment"` //[]Pending-payment;待支付切牌你
}


type Commidity_introduction struct {
	UID int `form:"UID"`
	Commidity_Name string `form:"Commidity_Name"`
	Volume int `form:"Volume"`//商品成交量
	Evaluations []string `form:"Evaluations"` //[]Evaluations;评论参数
	Detailed_Introduction string `form:"Detailed_Introduction"` //商品详细介绍
}


//JWT相关结构体
//声明claims结构体,自定义字段添加用户名
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
//定义secret
var Secret =[]byte("redholiday")
//定义Token有效时间
var JWT_Effective_Time = time.Minute*30
