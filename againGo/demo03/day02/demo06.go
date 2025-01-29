// @Author Gopher
// @Date 2025/1/29 14:28:00
// @Desc
package main

import "fmt"

func f1(a []int) {
	a[0] = 100
}

func main() {
	a := []int{1, 2}
	f1(a)
	fmt.Println(a)
}
