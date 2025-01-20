// @Author Gopher
// @Date 2025/1/19 20:35:00
// @Desc
package main

import "fmt"

func main() {

	// 创建一个无缓冲的Channel
	ch := make(chan int)

	// 创建一个有缓冲的Channel，缓冲大小为10
	//chBuffered := make(chan int, 10)

	// 向无缓冲的Channel写入数据
	//ch <- 1
	//ch <- 2
	//ch <- 3

	chBuffered := make(chan int, 2)
	chBuffered <- 1
	chBuffered <- 2
	chBuffered <- 3

	for v := range ch {
		fmt.Println(v)
	}

	for s := range chBuffered {
		fmt.Println(s)
	}

	close(ch)
	close(chBuffered)

}
