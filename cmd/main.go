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

	engine.GET("/redholiday/front-page",service.VerifyJWT(),api.FrontPage)


	engine.GET("/redholiday/:username",api.UserIntroduction)  //获取个人信息
	userGroup := engine.Group("/redholiday/user")
	{
		userGroup.POST("/regist",api.Regist)   //用户注册
		userGroup.POST("/login/pw",api.LoginPw)   //账号密码登录
		userGroup.POST("/login/email/verify",api.EmailLoginVerify)   //发送邮箱验证码
		userGroup.POST("/login/email",api.EmailLogin)   //邮箱验证码登录
		userGroup.GET("/shopping-cart", service.VerifyJWT(),api.ShoppingCart)   //查看购物车
		userGroup.POST("shopping-cart", service.VerifyJWT(),api.ShoppingCartRevise)  //更改购物车
		userGroup.GET("/order",service.VerifyJWT(),api.Order)  //查看订单
		userGroup.POST("/order",service.VerifyJWT(),api.OrderRevise)
		userGroup.GET("/login/github",)   //github登录接口
		userGroup.GET("/login/redirec",)   //github登录返回接口
	}

	commidityGroup := engine.Group("/redholiday/commidity")
	{
		commidityGroup.GET("/commiditys",api.Commiditys)  //查看所有商品
		commidityGroup.POST("/:uid",service.VerifyJWT(),api.CommidityIntroduction)  //商品详情及添加购物车
		commidityGroup.POST("/:uid", service.VerifyJWT(),api.CommidityComment)   //商品评论

	}

	engine.Run()

}