package main

import "fmt"

func main() {
	var number = make([]int, 3, 5)
	printSlice(number)
}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
