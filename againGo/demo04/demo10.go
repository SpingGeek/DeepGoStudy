// @Author Gopher
// @Date 2025/1/29 16:27:00
// @Desc
package main

import (
	"fmt"
	"runtime"
)

func show01(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}

func main() {
	go show01("java")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched() // 注释掉查看结果
		fmt.Println("golang")
	}
}
