package main

import "fmt"

func modifyValue(number int) {
	number = 11
}

func modifyValuePtr(number01 *int) {
	*number01 = 11
}
func main() {
	number := 7
	modifyValue(number)
	fmt.Println(number)
	number01 := 7
	modifyValuePtr(&number01)
	fmt.Println(number01)
}
