// @Author Gopher
// @Date 2025/1/20 10:02:00
// @Desc
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 10
	v := reflect.ValueOf(x) // 获取值的反射对象
	t := reflect.TypeOf(x)  // 获取类型的反射对象

	fmt.Println("Type:", t)  // Type: int
	fmt.Println("Value:", v) // Value: 10

	// 修改值（需要传递指针）
	p := reflect.ValueOf(&x)
	p.Elem().SetInt(20)
	fmt.Println("Updated Value:", x) // Updated Value: 20
}
