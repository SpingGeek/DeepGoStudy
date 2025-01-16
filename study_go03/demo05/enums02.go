package main

import "fmt"

func main() {

	const (
		B = 1 << (10 * iota)
		KB
		MB
		GB
	)

	fmt.Println(B, KB, MB, GB)
}
