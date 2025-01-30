// @Author Gopher
// @Date 2025/1/30 15:39:00
// @Desc
package main

import (
	"fmt"
	"sync"
)

var count int

func increment(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		count++
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()

	// 因为在两个协程下分别去去一个count的数目进行了相关的操作
	// 所以这个过程中的数量是随机的不安全的
	// 并不能去保证每一个执行的操作都可以的得到响应
	fmt.Println(count)
}
