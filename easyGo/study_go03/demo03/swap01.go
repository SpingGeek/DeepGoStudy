package main

import "fmt"

func main() {

	var a int = 7
	var b int = 11
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}
