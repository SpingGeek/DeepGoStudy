// @Author Gopher
// @Date 2025/1/29 11:04:00
// @Desc
package main

import (
	"fmt"
)

func main() {
	// 创建一个无缓冲的 channel
	ch := make(chan string)

	// 启动一个 goroutine，向 channel 中发送数据
	go func() {
		ch <- "apple"
		ch <- "banana"
		ch <- "cherry"
		close(ch) // 关闭 channel，表示不再发送数据
	}()

	// 使用 for range 遍历 channel 中的数据
	for fruit := range ch {
		fmt.Println(fruit)
	}
}
