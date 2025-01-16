package main

import "fmt"

func Demo() {

	// if the function has the many of defer, they are the way of LIFO to deal with
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}

func main() {
	Demo()
}
