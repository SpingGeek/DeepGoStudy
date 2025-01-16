package main

import "fmt"

func main() {

	score := [5]int{99, 96, 90, 85, 83}
	for i, v := range score {
		fmt.Printf(" 第 %d 名: %d 分\n\n", i+1, v)
	}
}
