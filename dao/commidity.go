package dao

import (
	"RedHoliday/model"
	"fmt"
	"time"
)

//通过uid查询商品详细信息
func QueryCommmidty(uid int)model.Commidity{
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var c model.Commidity
	querystr := "select * from commidity where uid=?"
	//查询信息
	r := db.QueryRow(querystr,uid).Scan(&c.Uid,&c.Commidity_name,&c.Volume,&c.Evaluations,&c.Detailed_Introduction)
	fmt.Println(r)
	return c
}


//查询所有商品信息
func QueryCommiditys()map[int]string{
	var commiditys = make(map[int]string,20)
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var commidity model.Commidity
	querystr := "select uid,name from commidity where uid >0"
	//查询信息
	rows, err := db.Query(querystr)
	if err != nil {
		fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
		return commiditys
	}
	for rows.Next() {
		err := rows.Scan(&commidity.Uid,&commidity.Commidity_name)
		if err != nil {
			fmt.Println(err.Error(),"  ",time.Now().Format("2006-01-02 15:04:05"))
			return commiditys
		}
		commiditys[commidity.Uid] = commidity.Commidity_name
	}
	return commiditys
}



//评论商品
func UpdateEcalutions(uid int,comment string){
	//连接数据库
	db := Link_mysql()

	updatestr := "update commidity set evaluations=? where uid = ?"
	_, err := db.Exec(updatestr, comment,uid)
	if err != nil {
		return
	}
}