// @Author Gopher
// @Date 2025/1/29 08:08:00
// @Desc
package main

import "fmt"

func generator() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	next := generator()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
