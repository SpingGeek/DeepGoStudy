package main

import "fmt"

func main() {
	const (
		identifier1 int = iota
		identifier2
		identifier3
	)

	fmt.Println(identifier1, identifier2, identifier3)
}
