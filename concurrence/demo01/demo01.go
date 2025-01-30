// @Author Gopher
// @Date 2025/1/30 14:52:00
// @Desc
package main

import (
	"fmt"
	"time"
)

// main 协程
// g1 协程
func main() {

	fmt.Printf("\"start\":%v\n", "start")

	// 匿名协程
	go func() {
		fmt.Printf("\"g1\":%v\n", "g1")
		time.Sleep(time.Second * 3)
	}()

	// 当主协程推出了之后，他的自协程也就不复存在了
	time.Sleep(time.Second * 3)

	fmt.Printf("\"end\":%v\n", "end")
}
