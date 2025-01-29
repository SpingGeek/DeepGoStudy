// @Author Gopher
// @Date 2025/1/29 11:45:00
// @Desc
package main

import "fmt"

func printBox() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	//var x, y int
	printBox()
}
