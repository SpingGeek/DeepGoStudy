package main

import "fmt"

func main() {

	var numbers []int
	var i int
	for i = 0; i < 10; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)
}
