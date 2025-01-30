// @Author Gopher
// @Date 2025/1/30 15:18:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func worker(stopCh chan bool) {
	for {
		select {
		default:
			fmt.Println("子协程执行中")
			time.Sleep(500 * time.Millisecond)
		case <-stopCh:
			fmt.Println("子协程被取消")
			return
		}
	}
}

func main() {
	stopCh := make(chan bool)
	go worker(stopCh)
	time.Sleep(2 * time.Second)
	stopCh <- true
	time.Sleep(1 * time.Second)
	fmt.Println("主协程执行完毕")
}
