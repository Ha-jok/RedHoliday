package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func main(){
	engine := gin.Default()

	engine.POST("/test/jwt", func(c *gin.Context) {
		//获取参数
		user := c.PostForm("user")

		//生成token
		tokenstring,_ := Generate_Token(user)

		c.JSON(http.StatusOK,gin.H{
			"code" : 2000,
			"message" : "success",
			"data" : gin.H{"token" : tokenstring},
		})
	})
	engine.GET("/test/jwt",Verify_token(), func(c *gin.Context) {
		username := c.MustGet("Username").(string)
		c.JSON(http.StatusOK,gin.H{
			"code" : 2000,
			"message" : "success",
			"data" : gin.H{"username" : username},
		})
	})

	engine.Run()
}








//定义claims结构体
type Claims struct {
	Username string `json:"Username`
	jwt.StandardClaims
}
//定义jwt过期时间，为五分钟
const JwtEffectiveTime = time.Hour*1
//定义secret
var Secret = []byte("redholiday")


func Generate_Token(username string)(string,error){
	//创建声明
	c := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JwtEffectiveTime).Unix(),
			Issuer: "redholiday-project",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(Secret)
}

//解析token
func Parse_Token(tokenstring string)(*Claims,error){
	//解析token
	token,err := jwt.ParseWithClaims(tokenstring,&Claims{},func(token *jwt.Token)(i interface{},err error){
		return Secret,nil
	})
	if err != nil {
		return nil, err
	}
	if claims,ok := token.Claims.(*Claims);ok && token.Valid{
		return claims,nil
	}
	return nil,errors.New("invalid token")
}

//jwt认证中间件
func Verify_token() func(c *gin.Context) {
	return func(c *gin.Context) {
		//接受请求时,jwt在authorization中

		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2003,
				"message" : "Authorization为空",
			})
			c.Abort()
			return
		}

		//按空格分割,提取出token
		JWT_information := strings.SplitN(authorization," ",2)
		if !(len(JWT_information) == 2 && JWT_information[0] == "Bearer") {
			c.JSON(http.StatusOK,gin.H{
				"msg" : "sdf",
			})
			c.Abort()
			return
		}

		claim,err := Parse_Token(JWT_information[1])
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2005,
				"msg" : "token无效",
			})
			c.Abort()
			return
		}

		//保存token中的username
		c.Set("Username",claim.Username)
		c.Next()

	}
}