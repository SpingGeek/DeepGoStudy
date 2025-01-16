package main

import (
	"fmt"
	_ "src/easyGo/study_go01/function/InitLib1"
	_ "src/easyGo/study_go01/function/InitLib2"
)

func init() {

	fmt.Println("libMain init")
}

func swap(x, y int) int {
	var temp int

	temp = x
	x = y
	y = temp

	return temp
}

/*
*
函数中引用传递是 *
变量中引用传递是 &
*/
func swapThePtr(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {

	fmt.Println("libMain main")

	var a int = 200
	var b int = 300

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	swap(a, b)

	fmt.Printf("交换后 a 的值为 : %d\n", a)
	fmt.Printf("交换后 b 的值为 : %d\n", b)

	fmt.Println("-------------------------")

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	swapThePtr(&a, &b)

	fmt.Printf("交换后 a 的值为 : %d\n", a)
	fmt.Printf("交换后 b 的值为 : %d\n", b)

}
