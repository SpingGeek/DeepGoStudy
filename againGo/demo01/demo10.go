// @Author Gopher
// @Date 2025/1/29 08:09:00
// @Desc
package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

func main() {

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(10))
	fmt.Println(triple(10))
}
