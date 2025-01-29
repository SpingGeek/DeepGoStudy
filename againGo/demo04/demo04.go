// @Author Gopher
// @Date 2025/1/29 15:16:00
// @Desc
package main

import "fmt"

func main() {
	var name string
	name = "tom"
	// p_name 指针类型
	var p_name *string
	// &name 取name地址
	p_name = &name
	fmt.Printf("name: %v\n", name)
	// 输出指针地址
	fmt.Printf("p_name: %v\n", p_name)
	// 输出指针指向的内容值
	fmt.Printf("*p_name: %v\n", *p_name)
}
