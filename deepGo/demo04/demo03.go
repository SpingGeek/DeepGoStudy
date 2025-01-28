// @Author Gopher
// @Date 2025/1/22 11:09:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(10)
	//close(ch)

	go func() {
		for v := range ch {
			fmt.Printf("v = %v\n", v)
		}
		fmt.Println(20)
	}()
	ch <- 4
	time.Sleep(2 * time.Second)
	ch <- 5
	fmt.Println("main")
}
