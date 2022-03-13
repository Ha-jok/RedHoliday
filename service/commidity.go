package service

import (
	"RedHoliday/dao"
	"RedHoliday/model"
	"fmt"
	"strconv"
)

//该go文件存放/commidity路径下的服务函数

//查询商品详细信息
func QueryCommidity(uid int)model.Commidity{

	//提取数据库信息
	commidity := dao.QueryCommmidty(uid)


	//返回信息
	return commidity

}


//添加购物车
func AddCart(uid int,username string){
	//提取用户信息
	user := QueryUserIntruduction(username)

	//更改用户购物车
	var cart string
	if user.Shopping_cart == "无"{
		cart = strconv.Itoa(uid)+","
	} else {
		cart = user.Shopping_cart+strconv.Itoa(uid)+","
	}
	//重新储存用户购物车
	dao.UpdateCart(username,cart)
}

//评论
func Comment(uid int,comment string){
	//提取商品信息
	commidity := QueryCommidity(uid)
	//更改评论字段
	var commentnew string
	if commidity.Evaluations == "无"{
		commentnew = comment+","
	} else {
		commentnew = commidity.Evaluations+comment+","
	}
	fmt.Println(commentnew)
	//重新储存商品评论
	dao.UpdateEcalutions(uid,commentnew)
}


//商品列表
func Commiditys()(map[int]string){
	//提取信息
	commiditys := dao.QueryCommiditys()
	return commiditys
}