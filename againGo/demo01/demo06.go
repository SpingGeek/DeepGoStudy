// @Author Gopher
// @Date 2025/1/28 22:05:00
// @Desc
package main

import "fmt"

func main() {
	i := 0

Loop:
	if i < 5 {
		fmt.Println("i = ", i)
		i++
		goto Loop
	}

	fmt.Println("循环结束")

}
