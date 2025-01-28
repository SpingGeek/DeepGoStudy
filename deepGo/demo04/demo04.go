// @Author Gopher
// @Date 2025/1/24 11:01:00
// @Desc
package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}
