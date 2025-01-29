// @Author Gopher
// @Date 2025/1/29 11:12:00
// @Desc
package main

import "fmt"

func f5() {
	// MY_LABEL:
	for i := 0; i < 5; i++ {
	MY_LABEL:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue MY_LABEL
			}
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}
}
func main() {
	f5()
}
