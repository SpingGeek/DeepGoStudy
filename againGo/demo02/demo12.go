// @Author Gopher
// @Date 2025/1/29 11:14:00
// @Desc
package main

import "fmt"

func f6() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				goto LABEL1
			}
			fmt.Println(i, j)
		}
	}
LABEL1:
	fmt.Println("label1")
}

func main() {
	f6()
}
