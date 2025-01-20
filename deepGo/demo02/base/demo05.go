// @Author Gopher
// @Date 2025/1/19 20:40:00
// @Desc
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	number := []int{1, 2, 3, 4, 5}
	sum := 0
	ch := make(chan int)

	for _, num := range number {
		go func(n int) {
			m.Lock()
			sum += n
			ch <- sum
			m.Unlock()
		}(num)
	}

	var finalSum int
	for range number {
		finalSum = <-ch
		fmt.Println("Current Sum: ", finalSum)
	}
	fmt.Println("Final Sum: ", finalSum)
}
