// Package client @Author Spring
// @Date 2025/1/3 21:34:00
// @Desc
package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立链接
	client, _ := rpc.Dial("tcp", "localhost:1234")

	var reply string
	err := client.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
