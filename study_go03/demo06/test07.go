package main

import "fmt"

//var slice = []int{2, 5, 8}
//slice01 := []int{1, 4, 7}

func main() {

	/*
		使用 make 函数初始化元素是 int 类型的切片 slice
		为了这个切片分配5个 int 类型的元素
		设置切片的容量为 10
	*/
	slice := make([]int, 5, 10)
	slice[0] = 10
	slice[1] = 5
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
}
