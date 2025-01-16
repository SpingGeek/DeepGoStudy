package main

import "fmt"

func main() {

	func(m int, n int) {
		fmt.Println(m + n)
	}(10, 20)
}
