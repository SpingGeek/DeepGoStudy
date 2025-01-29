// @Author Gopher
// @Date 2025/1/29 10:50:00
// @Desc
package main

import "fmt"

func f() {
	var a = [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i - this is a key: %d, v - this is a value: %v\n", i, v)
	}
}

func main() {
	f()
}
