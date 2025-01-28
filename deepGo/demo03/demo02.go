// @Author Gopher
// @Date 2025/1/21 14:38:00
// @Desc
package main

import "fmt"

func main() {

	array1 := [3]int{1, 2, 3}
	var array2 [3]int

	array2 = array1
	// update the value of array1
	array1[0] = 10
	fmt.Println(array1)
	fmt.Println(array2)
}
