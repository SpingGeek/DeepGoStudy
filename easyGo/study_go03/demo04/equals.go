package main

import "fmt"

func main() {
	var num01 = 6
	var num02 = 66
	var state01 bool = num01 == num02
	var state02 bool = num01 != num02
	fmt.Println(state01)
	fmt.Println(state02)
}
