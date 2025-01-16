package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 3
	n := 3
	numbers = append(numbers[:i], numbers[(i+n):]...)
	fmt.Println(numbers)
}
