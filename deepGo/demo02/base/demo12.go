// @Author Gopher
// @Date 2025/1/20 14:40:00
// @Desc
package main

import "fmt"

func counter(in chan<- int) {
	defer close(in)
	for i := 0; i < 100; i++ {
		in <- i
	}
}

func square(in chan<- int, out <-chan int) {
	defer close(in)
	for i := range out {
		in <- i * i
	}
}

func output(out <-chan int) {
	for i := range out {
		fmt.Println(i)
	}
}

// 改写成单向通道
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go square(ch2, ch1)
	output(ch2)
}
