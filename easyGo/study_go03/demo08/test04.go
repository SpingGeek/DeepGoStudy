package main

import "fmt"

func main() {

	str := "knowledge is power"
	ptr := &str
	fmt.Println(ptr)
	fmt.Printf("指针变量 ptr 的类型: %T\n", ptr)
	fmt.Printf("变量 str 的内存地址: %p\n", ptr)
	str_value := *ptr
	fmt.Printf("指针变量 str_value 的类型: %T\n", str_value)
	fmt.Printf("变量 str_value 的内存地址: %s\n", str_value)

}
