// @Author Gopher
// @Date 2025/1/21 15:17:00
// @Desc Go 示例：展示通过反射和 unsafe 操作打印切片的底层结构信息

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 创建一个长度为 5，容量为 10 的 int 类型切片
	s := make([]int, 5, 10)

	// 调用 PrintSlice 函数，传递切片的指针
	// 这里传入的是切片的地址，以便在 PrintSlice 函数中操作切片指针
	PrintSlice(&s)

	// 调用 test 函数，传入切片，展示函数间传递的切片信息
	test(s)
}

// test 函数接收切片并打印切片信息
func test(s []int) {
	// 调用 PrintSlice 函数，传递切片的指针
	PrintSlice(&s)
}

// PrintSlice 函数接收切片的指针，并通过反射和 unsafe 操作打印切片的底层数据结构信息
func PrintSlice(s *[]int) {
	// 使用 unsafe.Pointer 将切片指针转换为 reflect.SliceHeader 类型的指针
	// reflect.SliceHeader 结构体包含切片的底层数据结构（数组指针、长度、容量等）
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))

	// 打印切片的底层结构和切片本身的信息
	// ss：底层数据结构的信息
	// s：原始传入的切片指针
	fmt.Printf("slice struct : %+v, slice is %v\n", ss, s)
}
