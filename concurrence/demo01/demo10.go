// @Author Gopher
// @Date 2025/1/30 15:43:00
// @Desc
package main

import (
	"fmt"
	"sync"
)

var (
	count01 int
	m01     sync.Mutex
)

func increment01(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {

		m01.Lock()
		count01++
		m01.Unlock() // 注销就死锁
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment01(&wg)
	go increment01(&wg)
	wg.Wait()

	// 这个地方当我们在原本的方法执行的过程中加上锁
	// 这样的方法就可以保证在协程执行的过程中两个协程是不能去产生相互的影响的
	fmt.Println(count01)
}
