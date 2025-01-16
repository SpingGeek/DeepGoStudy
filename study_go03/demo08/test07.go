package main

import "fmt"

const length = 4

func main() {
	number := []int{2, 4, 6, 8}
	for i := 0; i < length; i++ {
		fmt.Printf("number[%d] = %d\n", i, number[i])
	}
}
