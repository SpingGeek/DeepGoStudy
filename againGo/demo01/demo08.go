// @Author Gopher
// @Date 2025/1/29 08:05:00
// @Desc
package main

import "fmt"

func modifyPointer(x *int) {
	*x = 100
	fmt.Println("Inside modifyPointer: ", *x)
}

func main() {
	a := 50
	fmt.Println("Before modifyPointer: ", a)
	modifyPointer(&a)
	fmt.Println("After modifyPointer: ", a)
}
