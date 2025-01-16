package main

import "fmt"

func main() {
	getNum := func(num int) bool {
		if num%3 == 0 && num%5 == 0 {
			return true
		} else {
			return false
		}
	}
	n := 0
	for i := 1; i <= 1000; i++ {
		if getNum(i) {
			n++
			fmt.Print(i, "\t\t")
			if n%6 == 0 {
				fmt.Print("\n")
			}
		}
	}
}
