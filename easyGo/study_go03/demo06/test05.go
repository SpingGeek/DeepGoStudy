package main

import "fmt"

func main() {

	var a = [4]int{1, 4, 3, 5}
	b := [4]int{2, 4, 6, 8}
	c := [...]int{1, 3, 5, 7}
	// 两个数组的长度不相同是不能去进行相关的比较的
	// 两个数组的类型不相同也不是进行比较的
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a == b)
	fmt.Println(b == c)
	fmt.Println(a != c)

}
