package service

import (
	"RedHoliday/dao"
	"RedHoliday/model"
	"fmt"
	"strconv"
)

//该go文件存放/commidity路径下的服务函数

//查询商品详细信息
func Query_commidity(uid int)model.Commidity{

	//提取数据库信息
	commidity := dao.Query_commmidty(uid)


	//返回信息
	return commidity

}


//添加购物车
func Add_cart(uid int,username string){
	//提取用户信息
	user := Query_user_intruduction(username)

	//更改用户购物车
	var cart string
	if user.Shopping_cart == "无"{
		cart = strconv.Itoa(uid)+","
	} else {
		cart = user.Shopping_cart+strconv.Itoa(uid)+","
	}
	//重新储存用户购物车
	dao.Update_cart(username,cart)
}

//评论
func Comment(uid int,comment string){
	//提取商品信息
	commidity := Query_commidity(uid)
	//更改评论字段
	var commentnew string
	if commidity.Evaluations == "无"{
		commentnew = comment+","
	} else {
		commentnew = commidity.Evaluations+comment+","
	}
	fmt.Println(commentnew)
	//重新储存商品评论
	dao.Update_ecalutions(uid,commentnew)
}
