// @Author Gopher
// @Date 2025/1/29 10:35:00
// @Desc
package main

import "fmt"

func main() {
	a := 4 // 二进制 100
	fmt.Printf("a: %b\n", a)
	b := 8 // 二进制 1000
	fmt.Printf("b: %b\n", b)

	fmt.Printf("(a & b): %v, %b \n", (a & b), (a & b))
	fmt.Printf("(a | b): %v, %b\n", (a | b), (a | b))
	fmt.Printf("(a ^ b): %v, %b\n", (a ^ b), (a ^ b))
	fmt.Printf("(a << 2): %v, %b\n", (a << 2), (a << 2))
	fmt.Printf("(b >> 2): %v, %b\n", (b >> 2), (b >> 2))
}
