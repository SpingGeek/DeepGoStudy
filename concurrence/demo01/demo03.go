// @Author Gopher
// @Date 2025/1/30 15:04:00
// @Desc
package main

import (
	"context"
	"fmt"
	"time"
)

// printNum 就是接受信息打印数字
func printNum(ctx context.Context, n int) {

	for i := 1; i <= n; i++ {
		// listener
		select {
		// 判断ctx.Done()是否关闭
		case <-ctx.Done():
			fmt.Println("协程被取消")
			return
		default:
			fmt.Printf("子协程中的数字：%d\n", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	// context.WithCancel 函数创建了一个带有取消功能的上下文对象 ctx
	ctx, cancel := context.WithCancel(context.Background())
	go printNum(ctx, 20)
	time.Sleep(6 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("主协程执行完毕")

}
