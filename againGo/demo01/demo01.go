// @Author Gopher
// @Date 2025/1/28 21:20:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func say(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func main() {
	go say("world")
	say("hello")
}
