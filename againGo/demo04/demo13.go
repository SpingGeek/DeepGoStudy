// @Author Gopher
// @Date 2025/1/29 16:44:00
// @Desc
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	counter := 1
	for _ = range ticker.C {
		fmt.Println("ticker 1")
		counter++
		if counter >= 5 {
			break
		}
	}
	ticker.Stop()
}
