package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Print("第", i, "次循环输出: ")
		for j := 1; j <= 5; j++ {
			if j == 3 {
				goto tag
			}
			fmt.Print(j, "")
		}
		fmt.Println()
	}
	fmt.Print("\n", "不会显示")
tag:
	fmt.Println("\nover")
}
