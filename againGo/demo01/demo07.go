// @Author Gopher
// @Date 2025/1/28 22:07:00
// @Desc
package main

import "fmt"

func main() {

	for i := 1; i < 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				goto EndLoop
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
EndLoop:
	fmt.Println("循环结束")
}
