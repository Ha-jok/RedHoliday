package api

//该go文件存放/commidity路径下接口，：商品详情,商品评论

import (
	"RedHoliday/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//返回商品列表,/redholiday/commiditys
func Commiditys (engine *gin.Engine){
	engine.GET("/redholiday/commiditys", func(c *gin.Context) {
		//获取map
		commiditys := service.Commiditys()
		c.JSON(http.StatusOK,gin.H{
			"data" : commiditys,
	})
})
}


//商品详情及添加购物车，/redholiday/commidity/:uid
func Commidity_introduction (engine *gin.Engine){
	engine.GET("/redholiday/commidity/:uid", service.Verify_JWT(),func(c *gin.Context) {
		//获取参数
		uid := c.Param("uid")
		add := c.PostForm("add")

		//查询信息
		//将uid转化为Int类型，
		uidm,_ := strconv.Atoi(uid)
		//提取信息
		commidity := service.Query_commidity(uidm)

		//返回信息
		c.JSON(http.StatusOK,gin.H{
			"data" : gin.H{
				"uid" : commidity.Uid,
				"name" : commidity.Commidity_name,
				"volume" : commidity.Volume,
				"evaluations" : commidity.Evaluations,
				"detailed" : commidity.Detailed_Introduction,
			},
		})

		//添加购物车
		if add == "add"{
			//提取token参数中的信息
			username := c.MustGet("username").(string)
			//更改数据库中的数据
			service.Add_cart(uidm,username)
		}
	})
}

//商品评论，/redholiday/commidity/:uid
func Commidity_comment (engine *gin.Engine){
	engine.POST("/redholiday/commidity/:uid", service.Verify_JWT(),func(c *gin.Context) {
		//获取参数
		Comment := c.PostForm("comment")
		uid := c.Param("uid")
		add := c.PostForm("add")
		//将uid转化为Int类型，
		uidm,_ := strconv.Atoi(uid)
		//评论
		if Comment != ""{
			//发送到数据库
			service.Comment(uidm,Comment)

			//返回信息
			c.JSON(http.StatusOK,gin.H{
				"message" : "评论成功",
			})
		}


		//添加购物车
		if add == "add"{
			//提取token参数中的信息
			username := c.MustGet("username").(string)
			//更改数据库中的数据
			service.Add_cart(uidm,username)
			c.JSON(http.StatusOK,gin.H{
				"message" : "添加成功",
			})
		}


	})
}


