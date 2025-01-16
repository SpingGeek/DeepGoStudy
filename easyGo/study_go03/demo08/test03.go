package main

import "fmt"

func main() {
	var num *int32
	num = new(int32)
	*num = 10
	fmt.Println(num)
}
