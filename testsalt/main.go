package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	time2 "time"
)

func main(){
	//获取当前时间，时分秒
	time := time2.Now().Format("15:04:05")
	fmt.Println(time)
	fmt.Printf("%T\n",time)
	fmt.Println(encry("ab"))
	fmt.Println(encry("ab"))





}
func encry(pw string)string{
	h := md5.New()
	h.Write([]byte(pw))
	return hex.EncodeToString(h.Sum(nil))

}
