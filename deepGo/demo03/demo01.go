// @Author Gopher
// @Date 2025/1/21 14:31:00
// @Desc
package main

import "fmt"

func main() {

	array := [3]int{1, 2, 3}
	// array delivery to the function
	test01(array)
	fmt.Printf("array outside: %p\n", &array)
}

func test01(array [3]int) {
	fmt.Printf("array inside: %p\n", &array)
}
