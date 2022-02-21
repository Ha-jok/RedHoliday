package main

import "fmt"

func main() {
	var a,b bool
	a = true
	b = false
	fmt.Println(a,b)
	if a {
		fmt.Println("a1")
	}
	if !a {
		fmt.Println("a2")
	}
	if b {
		fmt.Println("b1")
	}
	if !b {
		fmt.Println("b2")
	}
	var c bool
	fmt.Println(c)
}
