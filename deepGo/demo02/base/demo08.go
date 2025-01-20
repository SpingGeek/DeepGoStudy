// @Author Gopher
// @Date 2025/1/19 22:03:00
// @Desc
package main

import "fmt"

// 函数的定义
func f1(x int, y int) (ret int) {
	ret = x + y
	return x + y
}

// 无返回值的函数
func f2(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数也没有返回值的
func f3() {
	fmt.Println("1111")
}

// 没有参数 有返回值的
func f4() int {
	return 4
}

// 返回值可以命名也可以不命名
// 命名的返回值就相当于在函数中声明一个变量
func f5(x int, y int) (ret int) {
	ret = x + y //注意：因为已经在返回值中声明了ret，所以这里用= 而不是:= ,避免重复声明问题
	return      //因为已经在函数体中声明了ret，所以在return的时候不需要重复声明
}

// 多个返回值
func f6() (int, int) {
	return 1, 2
}

// 多个参数简写的方式
// 当参数的类型一致时，可以将连续的相同参数 前面参数的类型省略 比如：
func f7(x, y, z int, a, b string, c, d bool) int {
	return x + y + z
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f8(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

// defer 延迟执行
func deferDemo() {
	defer fmt.Println("111") //最先defer的语句最后执行
	defer fmt.Println("222")
	fmt.Println("333")
}

// go语言中函数没有默认参数的概念
func main() {
	r := f5(1, 2)
	fmt.Println(r)

	m, n := f6()
	fmt.Println(m, n)

	r7 := f7(1, 2, 3, "1", "1", true, false)
	fmt.Println(r7)

	f8("hah") //可变长度 不填也可以
	f8("hah", 1, 2, 3, 4)

	//延迟函数测试
	deferDemo()
}
