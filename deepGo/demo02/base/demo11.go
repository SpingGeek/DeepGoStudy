// @Author Gopher
// @Date 2025/1/20 14:37:00
// @Desc
package main

import "fmt"

func main() {

	// channel练习 go  for range从chan中取值
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine 把0-100写入到ch1通道中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine 从ch1中取值，值的平方赋值给 ch2
	go func() {
		for {
			i, ok := <-ch1 //通道取值后 再取值 ok = false
			if ok {
				ch2 <- i * i
			} else {
				break
			}
		}
		close(ch2)
	}()

	// 主goroutine 从ch2中取值 打印输出
	// for x := chan 有值取值，通道关闭时跳出goroutine
	for i := range ch2 {
		fmt.Println(i)
	}

}
