// @Author Gopher
// @Date 2025/1/19 21:40:00
// @Desc
package main

import "fmt"

func main() {
	i := 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	s := "hello王中阳"
	for i, i2 := range s {
		fmt.Printf("ikey值为:%v，i2字符为:%c\n", i, i2)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("打印结束")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("打印结束")

	flag := false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			fmt.Println("over")
			break
		}
	}

	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto xx //跳转到定义的label语句，即直接跳出了for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}

	// 定义label
xx:
	fmt.Println("over")
}
