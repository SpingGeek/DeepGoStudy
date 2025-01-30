// @Author Gopher
// @Date 2025/1/30 15:00:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func child() {

	for i := 1; i <= 5; i++ {

		fmt.Printf("child:%v\n", i)
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {

	go child()

	for i := 1; i <= 5; i++ {

		fmt.Printf("main:%v\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}
