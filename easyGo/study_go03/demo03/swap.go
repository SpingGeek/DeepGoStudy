package main

import "fmt"

func main() {

	//var a = 7
	//var b = 11
	//var c = 0

	var a int = 7
	var b int = 11
	var c int

	c = a
	a = b
	b = c

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
