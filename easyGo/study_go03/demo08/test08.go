package main

import "fmt"

const length01 = 4

func main() {
	number := []int{2, 4, 6, 8}
	var ptr [length01]*int
	fmt.Printf("%v,%T\n", ptr, ptr)
	for i := 0; i < length01; i++ {
		ptr[i] = &number[i]
	}
	fmt.Printf("%v,%T\n", ptr, ptr)
	for i := 0; i < length01; i++ {
		fmt.Printf("number[%d] = %d\n", i, ptr[i])
	}
}
