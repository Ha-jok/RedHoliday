package main

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

func main(){
	//发送邮件的邮箱
	usermail := "323150736@qq.com"
	//qq邮箱smt接口和地址
	smt := "smtp.qq.com"
	smtp_port := ":587"
	//授权码
	authorize := "nlfkdkycxypccabg"
	//头部信息
	auth := smtp.PlainAuth("",usermail,authorize,smt)
	//接收者邮箱
	to := []string{"1273447417@qq.com"}
	//发送者名字
	sender := "redholiday-project"
	user := usermail
	//邮件标题
	subject := "验证码"
	content_type := "Content-Type: text/plain;charset=UTF-8"

	//邮件内容
	//随机生成验证码，以当前时间戳为种子
	rand.Seed(time.Now().UnixNano())
	var verify_code string
	for i:=0;i<4;i++{
		ve := rand.Intn(10)
		verify_code = verify_code+strconv.Itoa(ve)

	}
	//???搞不懂，只知道是接口要求数据格式，直接用了
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + sender +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + verify_code)
	//调用接口，发送邮件
	err := smtp.SendMail(smt+smtp_port,auth,user,to,msg)
	if err != nil {
		fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
	}
}
