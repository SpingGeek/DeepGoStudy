// @Author Gopher
// @Date 2025/1/29 11:41:00
// @Desc
package main

import "fmt"

func main() {

	var num int

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			num += i
		}
	}

	fmt.Println(num)
}
