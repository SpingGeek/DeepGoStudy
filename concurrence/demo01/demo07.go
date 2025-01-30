// @Author Gopher
// @Date 2025/1/30 15:30:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("start")

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Done")
	}()

	fmt.Println("end")
	// 如果我注释掉这个最后的等待时间的话
	// 这个主线程并不会在意子线程的存在和执行
	// 反而会直接把相对应的线程直接停掉，这样的 “Done” 这段就不复存在了
	time.Sleep(2 * time.Second)

}
