// @Author Gopher
// @Date 2025/1/29 09:14:00
// @Desc
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	index := 2

	s = append(s[:index], s[index+1:]...)
	fmt.Println(s)

}
