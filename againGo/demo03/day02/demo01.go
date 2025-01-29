// @Author Gopher
// @Date 2025/1/29 11:38:00
// @Desc
package main

import "fmt"

func main() {

	var num int
	fmt.Println("请输入一个整数")
	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}
	if num%2 == 1 {
		fmt.Println("这是一个奇数")
	} else {
		fmt.Println("这是一个偶数")
	}
}
