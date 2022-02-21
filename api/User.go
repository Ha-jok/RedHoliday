package api

//该go文件包括所有/user下接口：新用户注册，账号密码登录，邮箱验证码登录，手机号找回密码，获取个人信息，查看购物车，查看用户订单，修改头像，

import (
	"RedHoliday/model"
	"RedHoliday/service"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//账号密码登录，redholiday/user/login/pw
func Login_pw(engine *gin.Engine){
	engine.POST("/redholiday/user/login/pw", func(c *gin.Context) {
		//绑定参数,用户结构体,处理错误
		var Person model.Person_login_pw
		err := c.ShouldBind(&Person)
		if err != nil{
			fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
			c.JSON(http.StatusOK,gin.H{
				"message" : "出现错误",
			})
			return
		}
		//参数不能为空
		if Person.UserName == ""{
			c.JSON(http.StatusOK,gin.H{
				"message" : "用户名不能为空",
			})
			return
		}
		if Person.Password == ""{
			c.JSON(http.StatusOK,gin.H{
				"message" : "密码不能为空",
			})
			return
		}
		//验证用户名和密码是否匹配，接受一个布尔值
		bool := service.Judge_up(Person.UserName,Person.Password)
		//如果接收布尔值为false
		if !bool {
			c.JSON(http.StatusOK,gin.H{
				"message" : "用户名或密码错误",
			})
			return
		}
		//bool为true
		//在后台打印日志
		fmt.Println(Person.UserName,"登录成功")
		//返回token
		tokenstring,_ := service.Generate_Token(Person.UserName)
		c.JSON(http.StatusOK,gin.H{
			"code" : 2000,
			"message" : "欢迎回来，"+Person.UserName,
			"data" : gin.H{
				"token" : tokenstring,
			},
		})
	})
}



//新用户注册，/redholiday/user/regist
func Regist(engine *gin.Engine){
	engine.POST("/redholiday/user/regist", func(c *gin.Context) {
		//绑定参数并处理错误
		var New_Person model.New_person
		err := c.ShouldBind(&New_Person)
		if err != nil{
			fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
			c.JSON(http.StatusOK,gin.H{
				"message" : "注册失败，请检查输入信息或联系管理员",
			})
			return
		}
		//检验参数是否合法
		//验证用户名
		if len(New_Person.Username)>6||len(New_Person.Username)==0{
			c.JSON(http.StatusOK,gin.H{
				"message" : "用户名格式错误",
			})
			return
		}
		//验证密码
		if len(New_Person.Password)<6||len(New_Person.Password)>12{
			c.JSON(http.StatusOK,gin.H{
				"message" : "密码格式错误",
			})
			return
		}
		//验证电话号码
		if len(strconv.Itoa(New_Person.Phone)) != 11{
			c.JSON(http.StatusOK,gin.H{
				"message" : "电话号码格式错误",
			})
			return
		}
		//验证邮箱
		b := service.VerifyEmailFormat(New_Person.Email)
		if !b {
			c.JSON(http.StatusOK,gin.H{
				"message" : "邮箱格式不对",
			})
			return
		}
		//发送验证码



		//检验验证码是否正确



		//储存到数据库,添加到用户总表，并创建一个数据表
		//获取盐值,储存加密后的密码和盐值
		salt := service.Create_salt()
		password := New_Person.Password+salt
		password1 := service.Encry_pw(password)
		b = service.Create_user(New_Person.Username,password1,New_Person.Email,salt,New_Person.Phone)

		//返回数据
		if b {
			fmt.Println(New_Person.Username,"注册成功")
			c.JSON(http.StatusOK,gin.H{
				"message" : "注册成功"+New_Person.Username,
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"message" : "用户名重复，请更换用户名重新注册",
		})


	})

}

//手机验证码登录，/redholiday/user/login/phone
//未经过调试
func Login_phone(engine *gin.Engine){
	engine.POST("/redholiday/user/login/email", func(c *gin.Context) {
		//绑定参数并处理错误
		var Person model.Person_login_email
		err := c.ShouldBind(&Person)
		if err != nil{
			fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
			c.JSON(http.StatusOK,gin.H{
				"message" : "出现错误",
			})
			return
		}
		//验证手机号



		//发送验证码



		//检验验证码是否正确


		//从数据库中提取用户名
		var username string


		//返回数据及token
		tokenstring,err := service.Generate_Token(username)
		fmt.Println(username,"登录成功")
		c.JSON(http.StatusOK,gin.H{
			"code" : 2000,
			"data" : gin.H{
				"token" : tokenstring,
			},
		})

	})
}


//邮箱找回密码,/redholiday/user/forget-password
//发送验证码未完成
func Forget_Password(engine *gin.Engine){
	engine.POST("/redholiday/user/forget-password", func(c *gin.Context) {
		//获取参数并处理错误
		var Person model.Person_forget_password
		err := c.ShouldBind(&Person)
		if err != nil{
			fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
			c.JSON(http.StatusOK,gin.H{
				"message" : "出现错误",
			})
			return
		}
		//验证新密码是否符合格式
		if len(Person.New_password)<6||len(Person.New_password)>12{
			c.JSON(http.StatusOK,gin.H{
				"message" : "密码格式错误",
			})
			return
		}
		//发送验证码


		//检验验证码


		//储存数据库


		//返回信息


	})
}


//获取个人信息,/redholiday/user/:uid
func User_introduction(engine *gin.Engine){
	engine.GET("/redholiday/user/:username", func(c *gin.Context) {
		//获取username
		username := c.Param("username")


		//从数据库中提取信息
		user := service.Query_user_intruduction(username)

		//判断用户名是否有效
		if user.Username == "" {
			c.JSON(http.StatusOK,gin.H{
				"message" : username+"无效",
			})
		}




		//返回信息
		c.JSON(http.StatusOK,gin.H{

			"uid" : user.Uid,
			"username" : user.Username,
			"friends" : user.Friends,
			"follow_business" : user.Follow_business,

		})



	})
}


//查看购物车,/redholiday/user/shopping-cart
func Shopping_cart(engine *gin.Engine){
	engine.GET("/redholiday/user/shopping-cart", service.Verify_JWT(),func(c *gin.Context) {
		//提取token中的信息
		username := c.MustGet("username").(string)
		//提取数据库信息
		user := service.Query_user_intruduction(username)
		shopping_cart := user.Shopping_cart
		//将信息转换为切片
		shopping := strings.Split(shopping_cart,",")
		//返回信息
		c.JSON(http.StatusOK,gin.H{
			"code" : 2000,
			"data" : gin.H{
				"Shopping-cart" : shopping,
			},
		})
	})
}

//更改购物车,/redholiday/user/shopping-cart
func Shopping_cart_revise(engine *gin.Engine){
	engine.POST("/redholiday/user/shopping-cart", service.Verify_JWT(),func(c *gin.Context){
		//提取token中的信息
		username := c.MustGet("username").(string)
		//绑定参数
		settlement := c.PostForm("settlement")
		delete := c.PostForm("delete")
		//删除商品
		if delete != ""{
			//提取用户信息
			user := service.Query_user_intruduction(username)
			shopping_carts := user.Shopping_cart
			if shopping_carts == "无" {
				return
			}
			deletes := strings.Split(delete,",")
			//更改用户信息
			for _,v := range deletes{
				v1 := v+","
				shopping_carts = strings.Replace(shopping_carts,v1,"",1)
			}
			//重新储存用户信息
			service.Shopping_cart_revise(username,shopping_carts)
			//返回信息
			shopping_cart := strings.Split(shopping_carts,",")
			c.JSON(http.StatusOK,gin.H{
				"message" : "删除成功",
				"data" : gin.H{
					"shopping-cart" : shopping_cart,
				},
			})
		}
		//结算商品
		if settlement != "" {
			//提取用户信息
			user := service.Query_user_intruduction(username)
			shopping_carts := user.Shopping_cart
			if shopping_carts == "无" {
				return
			}
			settlements := strings.Split(delete,",")
			//更改用户信息
			for _,v := range settlements {
				v1 := v+","
				shopping_carts = strings.Replace(shopping_carts,v1,"",1)
			}
			//重新储存用户信息，购物车和待支付订单
			service.Shopping_cart_revise(username,shopping_carts)
			service.Order_paid(username,settlement)
			//返回信息
			shopping_cart := strings.Split(shopping_carts,",")
			c.JSON(http.StatusOK,gin.H{
				"message" : "结算成功",
				"data" : gin.H{
					"shopping-cart" : shopping_cart,
				},
			})
		}
	})
}

//查看用户订单,/redholiday/user/order
func Order(engine *gin.Engine){
	engine.GET("/redholiday/user/order",service.Verify_JWT(), func(c *gin.Context) {
		//验证token
		username := c.MustGet("username").(string)
		fmt.Println(username)
		//提取数据库信息
		user := service.Query_user_intruduction(username)
		order_p := user.Order_paid
		order_un := user.Order_unpaid
		order_re := user.Order_received
		//解析信息
		order_paid := strings.Split(order_p,",")
		order_unpaid := strings.Split(order_un,",")
		order_received := strings.Split(order_re,",")
		//返回信息
		c.JSON(http.StatusOK,gin.H{
			"data" : gin.H{
				"Order_paid" : order_paid,
				"Order_unpaid" : order_unpaid,
				"Order_received" : order_received,
			},
		})




	})
}


//修改用户订单，确认收货或取消订单,/redholiday/user/order
func Order_revise(engine *gin.Engine){
	engine.POST("/redholiday/user/order",service.Verify_JWT(), func(c *gin.Context) {
		//验证token
		username := c.MustGet("username").(string)
		fmt.Println(username)
		//绑定参数
		receit := c.PostForm("receit")
		cancel := c.PostForm("cancel")
		//提取数据库信息
		user := service.Query_user_intruduction(username)
		order_p := user.Order_paid
		order_un := user.Order_unpaid
		order_re := user.Order_received
		//取消订单
		if cancel != "" {
			if order_un == "无" {
				return
			}
			cancels := strings.Split(cancel,",")
			//更改用户信息
			for _,v := range cancels{
				v1 := v+","
				order_un = strings.Replace(order_un,v1,"",1)
			}
			//重新储存用户信息
			service.Order_unpaid(username,order_un)
			//解析信息
			order_paid := strings.Split(order_p,",")
			order_unpaid := strings.Split(order_un,",")
			order_received := strings.Split(order_re,",")
			//返回信息
			c.JSON(http.StatusOK,gin.H{
				"data" : gin.H{
					"Order_paid" : order_paid,
					"Order_unpaid" : order_unpaid,
					"Order_received" : order_received,
				},
			})
		}
		//确认收获
		if receit != "" {
			if order_p == "无" {
				return
			}
			if strings.Count(order_un,cancel) == 0{
				return
			}
			receits := strings.Split(receit,",")
			//更改用户信息
			for _,v := range receits {
				v1 := v+","
				order_p = strings.Replace(order_p,v1,"",1)
			}
			//重新储存用户信息
			order_re = order_re+receit+","
			service.Order_unpaid(username,order_un)
			service.Order_received(username,order_re)

			//解析信息
			order_paid := strings.Split(order_p,",")
			order_unpaid := strings.Split(order_un,",")
			order_received := strings.Split(order_re,",")
			//返回信息
			c.JSON(http.StatusOK,gin.H{
				"data" : gin.H{
					"Order_paid" : order_paid,
					"Order_unpaid" : order_unpaid,
					"Order_received" : order_received,
				},
			})
		}
	})
}