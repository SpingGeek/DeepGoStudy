package main

import "fmt"

func main() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", j*i, "\t")
		}
		fmt.Printf("\n")
	}
}
