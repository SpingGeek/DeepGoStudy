// @Author Gopher
// @Date 2025/1/19 17:02:00
// @Desc
package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 6)
	slice2 := append(slice1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
