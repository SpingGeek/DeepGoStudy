// @Author Gopher
// @Date 2025/1/29 14:51:00
// @Desc
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {

	add := cal("+")
	r := add(1, 2)
	fmt.Println(r)
}
