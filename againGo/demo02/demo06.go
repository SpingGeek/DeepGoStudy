// @Author Gopher
// @Date 2025/1/29 10:54:00
// @Desc
package main

import "fmt"

func f1() {
	var s = "多课网,go教程"
	for i, v := range s {
		fmt.Printf("i: %d, v: %c\n", i, v)
	}
	// %c 按照字符输出
}

func main() {
	f1()
}
