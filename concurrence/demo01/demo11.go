// @Author Gopher
// @Date 2025/1/30 15:50:00
// @Desc
package main

import (
	"fmt"
	"sync"
)

var x = 0
var m sync.Mutex

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		for i := 0; i < 1000000; i++ {
			m.Lock()
			x++
			m.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			m.Lock()
			x--
			m.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(x)
}
