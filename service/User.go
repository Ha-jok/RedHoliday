package service

import (
	"RedHoliday/dao"
	"RedHoliday/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"net/http"
	"net/smtp"
	"regexp"
	"strconv"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"errors"
)

//该go文件存放/user路径下所调用的非jwt服务函数





//判断密码是否正确
func JudgeUp(un,pw string)bool{
	//提取盐值和加密后的账号
	upw,salt := dao.QueryUp(un)
	//对用户输入密码进行加密
	pw1 := pw+salt
	pwj := EncryPw(pw1)
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
func CreateUser(username,password,email,salt string,phone int)bool{
	//将参数传入数据库,储存到用户表
	b := dao.InsertNewPerson(username,password,email,salt,phone)
	if !b {
		return false
	}
	//获取用户名
	uid := dao.QueryUid(username)
	//为用户创建一个个人数据表
	b = dao.CreateNewPerson(username)
	if !b {
		return false
	}
	//将信息储存到新表中
	b = dao.InsertNewTable(username,uid)
	if !b {
		return false
	}
	return true

}

//根据username查询个人信息（购物车，订单不可查询）
func QueryUserIntruduction(username string)model.Person_mysql{


	//从数据库中提取信息
	user := dao.QueryUsernameIntroduction(username)


	//将信息返回
	return user


}


//注册用户时获取盐值
func CreateSalt()string{
	//获取当前时间作为盐值
	salt := time.Now().Format("15:04:05")

	//返回时间戳
	return salt

}


//使用md5加密用户密码
func EncryPw(pw string)string{
	h := md5.New()
	h.Write([]byte(pw))
	return hex.EncodeToString(h.Sum(nil))
}


//修改用户购物车商品,重新储存购物车信息
func ShoppingCartRevise(username,shopping_carts string){
	//重新储存用户购物车
	dao.UpdateCart(username,shopping_carts)
}

//储存支付订单
func OrderPaid(username,settlement string){
	//提取原有支付订单
	user := QueryUserIntruduction(username)
	paid := user.Order_paid+settlement+","
	//更改用户订单状态
	dao.UpdateOrderPaid(username,paid)
}

//修改待支付订单
func OrderUnpaid(username,order_un string){
	//更改用户订单状态
	dao.UpdateOrderUnpaid(username,order_un)
}

//修改确认收货订单
func OrderReceived(username,receit string){
	dao.UpdateOrderReceived(username,receit)
}


//发送邮箱验证码
func EmailVerifyCode(email string)string{
	//绑定邮箱的地址,发送验证码的邮箱
	sender_email := "323150736@qq.com"
	//设置smtp，qq邮箱地址及端口，授权码
	smt := "smtp.qq.com"
	smtp_port := ":587"
	authorize_password := "nlfkdkycxypccabg"
	//头部信息
	auth := smtp.PlainAuth("",sender_email,authorize_password,smt)
	//设置邮件发送内容类型
	content_type := "Content-Type: text/plain;charset=UTF-8"
	//转变收件人邮箱格式为切片
	receiver := []string{email}
	//设置发送信息，发件人，标题，内容
	//发件人
	sender_name := "redholiday-project"
	//标题
	title := "验证码"
	//内容，随机四位数的验证码，利用时间戳和rand随机数
	rand.Seed(time.Now().UnixNano())
	var verify_code string
	for i:=0;i<4;i++{
		ve := rand.Intn(10)
		verify_code = verify_code+strconv.Itoa(ve)

	}
	//配置发送信息格式
	msg := []byte("To: " + strings.Join(receiver, ",") + "\r\nFrom: " + sender_name +
		"<" + sender_email + ">\r\nSubject: " + title + "\r\n" + content_type + "\r\n\r\n" + verify_code)
	//调用接口发送信息,并处理错误
	err := smtp.SendMail(smt+smtp_port,auth,sender_email,receiver,msg)
	if err != nil {
		fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
		return " "
	}
	return verify_code

}

//通过邮箱提取用户密码
func QueryEmailPw(email string)string{
	//提取数据
	username := dao.QueryEmail(email)
	return username
}

//jwt生成
func GenerateToken(username string)(string,error){
	//创建声明
	c := model.Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(model.JWT_Effective_Time).Unix(),
			Issuer: "redholiday-project",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(model.Secret)

}


//JWT解析
func ParseJWT(tokenstring string)(*model.Claims,error){
	//解析token
	token,err := jwt.ParseWithClaims(tokenstring,&model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return model.Secret,nil
	})
	//处理错误
	if err != nil {
		return nil, err
	}
	//验证是否token有效
	if claims,ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims,nil
	}
	return nil, errors.New("token无效")
}


//JWT认证
func VerifyJWT()func(c *gin.Context){
	return func(c *gin.Context) {

		//获取含有token信息的头部Authorazition部分
		authorization := c.Request.Header.Get("Authorization")
		fmt.Println(authorization)
		if authorization == "" {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2003,
				"message" : "Authorazition为空",
			})
			c.Abort()
			return
		}

		//提取token信息段
		JWT_information := strings.SplitN(authorization," ",2)
		//验证auth信息段是否合法
		if !(len(JWT_information) == 2 && JWT_information[0] == "Bearer") {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2004,
				"message" : "Authorazition格式错误",
			})
			c.Abort()
			return
		}
		//验证token是否有效
		claim,err := ParseJWT(JWT_information[1])
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2005,
				"message" : "token无效",
			})
			c.Abort()
			return
		}
		//将claim信息保存到上下文
		c.Set("username",claim.Username)
		c.Next()
	}
}