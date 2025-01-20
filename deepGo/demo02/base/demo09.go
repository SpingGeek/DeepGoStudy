// @Author Gopher
// @Date 2025/1/19 22:06:00
// @Desc
package main

import "fmt"

func f20() int {
	return 2
}

// 函数也可以作为函数参数的类型
func f30(x func() int) {
	ret := x()
	fmt.Printf("f3打印ret的值：%v\n", ret)  //2
	fmt.Printf("f3打印ret的类型：%T\n", ret) //int
}

func main() {
	a := f20
	fmt.Printf("a的类型：%T\n", a)
	f30(a)
}
