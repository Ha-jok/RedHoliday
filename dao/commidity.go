package dao

import (
	"RedHoliday/model"
)

//通过uid查询商品详细信息
func QueryCommmidty(uid int)model.Commidity{
	//连接数据库
	db := Link_mysql()
	var c model.Commidity
	//查询信息
	db.Table("commidity").Where("uid = ?",uid).Find(&c)
	return c
}


//查询所有商品信息
func QueryCommiditys()map[int]string{
	var commiditys = make(map[int]string,20)
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var commidity model.Commidity
	//查询信息
	db.Table("commidity").Where("uid > ?",0).Find(&commidity)
		commiditys[commidity.Uid] = commidity.Commidity_name
	return commiditys
}



//评论商品
func UpdateEcalutions(uid int,comment string){
	//连接数据库
	db := Link_mysql()
	db.Table("commidity").Where("uid = ?",uid).Update("evaluations",comment)
}