package main

import "fmt"

func main() {

	fmt.Println("begin")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
