// @Author Gopher
// @Date 2025/1/28 21:52:00
// @Desc
package main

import "fmt"

func main() {

	var arr [5]int
	arr[0] = 10
	arr[1] = 20

	var arr2 = [3]string{"apple", "banana", "cherry"}

	arr3 := [...]float64{1.1, 2.2, 3.3}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)

	for i, v := range arr2 {

		fmt.Printf("Index: %d, Value: %s\n", i, v)

	}

	var s []int

	s1 := make([]int, 5)
	s2 := make([]int, 5, 10)

	s3 := []string{"Go", "Java", "Python"}

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
