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




//新用户注册
func Regist(c *gin.Context){
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
	//储存到数据库,添加到用户总表，并创建一个数据表
	//获取盐值,储存加密后的密码和盐值
	salt := service.CreateSalt()
	password := New_Person.Password+salt
	password1 := service.EncryPw(password)
	b = service.CreateUser(New_Person.Username,password1,New_Person.Email,salt,New_Person.Phone)

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
}

//账号密码登录
func LoginPw(c *gin.Context){
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
	bool := service.JudgeUp(Person.UserName,Person.Password)
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
	tokenstring,_ := service.GenerateToken(Person.UserName)
	c.JSON(http.StatusOK,gin.H{
		"code" : 2000,
		"message" : "欢迎回来，"+Person.UserName,
		"data" : gin.H{
			"token" : tokenstring,
		},
	})
}








var verifyCodeSend,username string
//接口一，发送验证码
func EmailLoginVerify (c *gin.Context){
	email := c.PostForm("email")
	//判断邮箱格式
	b := service.VerifyEmailFormat(email)
	if !b {
		c.JSON(http.StatusOK,gin.H{
			"message" : "邮箱格式错误",
		})
		return
	}
	//邮箱是否存在
	//提取信息
	username = service.QueryEmailPw(email)
	if username == "" {
		c.JSON(http.StatusOK,gin.H{
			"message" : "用户不存在",
		})
		return
	}
	//发送验证码
	verifyCodeSend= service.EmailVerifyCode(email)
}
//接口二，验证验证码是否正确
func EmailLogin (c *gin.Context){
	verify_code := c.PostForm("verify_code")
	//获取上下文中的验证码
	if verify_code== "" {
		c.JSON(http.StatusOK,gin.H{
			"message" : "验证码不能为空",
		})
		return
	}
	if verifyCodeSend == verify_code{
		//在后台打印日志
		fmt.Println(username,"登录成功")
		//返回token
		tokenstring,_ := service.GenerateToken(username)
		c.JSON(http.StatusOK,gin.H{
			"code" : 2000,
			"message" : "欢迎回来，"+username,
			"data" : gin.H{
				"token" : tokenstring,
			},
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message" : "验证码错误",
	})
}





//获取个人信息
func UserIntroduction (c *gin.Context){
	//获取username
	username := c.Param("username")


	//从数据库中提取信息
	user := service.QueryUserIntruduction(username)

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
}


//查看购物车
func ShoppingCart(c *gin.Context){
	//提取token中的信息
	username := c.MustGet("username").(string)
	//提取数据库信息
	user := service.QueryUserIntruduction(username)
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
}

//更改购物车,/redholiday/user/shopping-cart
func ShoppingCartRevise(c *gin.Context){
	//提取token中的信息
	username := c.MustGet("username").(string)
	//绑定参数
	settlement := c.PostForm("settlement")
	delete := c.PostForm("delete")
	//删除商品
	if delete != ""{
		//提取用户信息
		user := service.QueryUserIntruduction(username)
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
		service.ShoppingCartRevise(username,shopping_carts)
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
		user := service.QueryUserIntruduction(username)
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
		service.ShoppingCartRevise(username,shopping_carts)
		service.OrderPaid(username,settlement)
		//返回信息
		shopping_cart := strings.Split(shopping_carts,",")
		c.JSON(http.StatusOK,gin.H{
			"message" : "结算成功",
			"data" : gin.H{
				"shopping-cart" : shopping_cart,
			},
		})
	}
}

//查看用户订单,/redholiday/user/order
func Order(c *gin.Context){
	//验证token
	username := c.MustGet("username").(string)
	fmt.Println(username)
	//提取数据库信息
	user := service.QueryUserIntruduction(username)
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
}


//修改用户订单，确认收货或取消订单,/redholiday/user/order
func OrderRevise(c *gin.Context){
	//验证token
	username := c.MustGet("username").(string)
	fmt.Println(username)
	//绑定参数
	receit := c.PostForm("receit")
	cancel := c.PostForm("cancel")
	//提取数据库信息
	user := service.QueryUserIntruduction(username)
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
		service.OrderUnpaid(username,order_un)
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
		service.OrderUnpaid(username,order_un)
		service.OrderReceived(username,order_re)

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
}