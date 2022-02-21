package main

import (
	"fmt"
	"strings"
)

func main(){
	a := "a,b,c,d,f,g,"
	b := strings.Split(a,",")
	fmt.Println(b)
	fmt.Println(len(b))
	a = strings.Replace(a,"a,","",-1)
	fmt.Println(a)

}
