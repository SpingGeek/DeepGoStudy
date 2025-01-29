// @Author Gopher
// @Date 2025/1/29 11:52:00
// @Desc
package main

import (
	"fmt"
	"math"
)

func checkPrimeNumber(num int) {
	if num < 2 {
		fmt.Println("不是素数")
		return
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			fmt.Println("不是素数")
			return
		}
	}
	fmt.Println("是素数")
}
func main() {
	var num int
	fmt.Println("请输入一个整数")
	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}
	checkPrimeNumber(num)
}
