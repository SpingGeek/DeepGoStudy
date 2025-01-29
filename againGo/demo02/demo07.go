// @Author Gopher
// @Date 2025/1/29 11:02:00
// @Desc
package main

import "fmt"

func f2() {
	var s = []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Printf("i, %d, v: %v\n", i, v)
	}
}

func main() {
	f2()
}
