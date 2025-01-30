// @Author Gopher
// @Date 2025/1/30 15:33:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("New: 协程已被创建但还未开始执行任务")

}

func bar() {
	fmt.Println("Runnable: 协程正在执行任务")
	time.Sleep(time.Second)
}

func baz() {
	fmt.Println("Blocked: 协程因为等待某些条件而被暂停执行")
	time.Sleep(time.Second)
}

func main() {

	go foo()

	go bar()

	go baz()

	// 阻塞状态
	ch := make(chan bool)
	go func() {
		fmt.Println("Blocked: 协程因为等待channel接收数据而被暂停执行")
		<-ch
	}()

	// 死亡状态
	go func() {
		fmt.Println("Dead: 协程执行完成")
	}()

	time.Sleep(time.Second * 2)
}
