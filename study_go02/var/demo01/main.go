package main

import "fmt"

// global variable
var n7 = 100

var n8 = 9.7

var (
	n9  = 500
	n10 = "netty"
)

func main() {

	var age int
	age = 18
	fmt.Println(age)

	var age01 int = 19
	fmt.Println(age01)

	/*
		when define the same as the name of var is error
		when define the var , but not to use , is error
	*/

	var name = "Tom"
	fmt.Println(name)

	sex := "man"
	fmt.Println(sex)

	fmt.Println(n7)
	fmt.Println(n8)
	fmt.Println(n9)
	fmt.Println(n10)

	fmt.Println(nil)
}
