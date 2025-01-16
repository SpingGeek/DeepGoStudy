package main

import "fmt"

func main() {
	number := 10
	str := "success"
	fmt.Printf("变量 number 的内存地址是 %p\n", &number)
	fmt.Printf("变量 str 的内存地址是 %p", &str)
}
