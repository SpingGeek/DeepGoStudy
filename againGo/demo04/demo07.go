// @Author Gopher
// @Date 2025/1/29 16:10:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func show(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go show("java")
	show("golang")
	fmt.Println("end...")
}
