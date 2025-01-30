// @Author Gopher
// @Date 2025/1/30 15:27:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	//time.Sleep(1 * time.Second)
	go func() {
		fmt.Println("hello world")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}
