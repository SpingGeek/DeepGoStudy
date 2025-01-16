package main

import "fmt"

func first(num int) int {
	fmt.Println("I am first")
	return num + 10
}

func second(num int) int {
	fmt.Println("I am second")
	return num + 20
}

func third(num int) int {
	fmt.Println("I am third")
	return num + 30
}

func main() {
	num := 10
	fmt.Println("begin")
	defer first(num)
	defer fmt.Println(second(num))
	num = 20
	defer fmt.Println(third(num))
	fmt.Println("end")
}

// 192
