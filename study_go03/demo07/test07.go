package main

import "fmt"

func main() {

	sum := func(m, n int) int {
		return m + n
	}
	result1 := sum(10, 20)
	result2 := sum(100, 200)
	fmt.Println(result1)
	fmt.Println(result2)
}
