package main

import (
	"RedHoliday/api"
	"RedHoliday/service"
	"github.com/gin-gonic/gin"
)

func main(){
	engine := gin.Default()

	//解决跨域
	engine.Use(service.Cross())

	//创建主页路由
	//测试成功,游客模式和登录模式成功判断
	api.Front_page(engine)

	//新用户注册，/redholiday/user/regist
	//测试成功
	api.Regist(engine)

	//账号密码登录，redholiday/user/login/pw
	//测试成功，token有效
	api.Login_pw(engine)

	//手机验证码登录，/redholiday/user/login/email
	api.Login_phone(engine)

	//邮箱找回密码,/redholiday/user/forget-password
	api.Forget_Password(engine)

	//获取个人信息,
	//测试完成
	api.User_introduction(engine)

	//查看购物车,
	//测试完成
	api.Shopping_cart(engine)

	//查看用户订单,
	//测试完成
	api.Order(engine)

	//商品详情，/redholiday/commidity/:uid
	//测试成功
	api.Commidity_introduction(engine)

	//商品评论及添加购物车，/redholiday/commidity/:uid
	//测试完成
	api.Commidity_comment(engine)

	//更改购物车,/redholiday/user/shopping-cart
	api.Shopping_cart_revise(engine)


	//修改用户订单，确认收货或取消订单,/redholiday/user/order
	api.Order_revise(engine)


	engine.Run()



}