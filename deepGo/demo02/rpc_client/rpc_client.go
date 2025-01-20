// @Author Gopher
// @Date 2025/1/20 14:47:00
// @Desc
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// ArithRequest 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// ArithResponse 响应结构体
type ArithResponse struct {
	Pro int //乘
	Quo int //商
	Rem int //余数
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("dialing error:", err)
	}

	req := ArithRequest{10, 20}
	var res ArithResponse

	err = conn.Call("Arith.Multiply", req, &res) //乘法运算
	if err != nil {
		log.Fatalln("arith error:", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	//除法运算
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("arith error:", err)
	}
	fmt.Printf("%d / %d = %d 余数是:%d", req.A, req.B, res.Quo, res.Rem)
}
