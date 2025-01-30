// @Author Gopher
// @Date 2025/1/30 15:14:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func printNum01(stopCh <-chan struct{}, n int) {

	for i := 1; i <= n; i++ {
		select {
		case <-stopCh:
			fmt.Println("协程被取消")
			return
		default:
			fmt.Printf("子协程中的数字:%d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	stopCh := make(chan struct{}, 10)
	go printNum01(stopCh, 10)

	time.Sleep(3 * time.Second)
	close(stopCh)

	time.Sleep(1 * time.Second)
	fmt.Println("主协程执行完毕")
}
