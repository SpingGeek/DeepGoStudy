// @Author Gopher
// @Date 2025/1/20 14:47:00
// @Desc
package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Arith struct {
}

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	rpc.Register(new(Arith))
	rpc.HandleHTTP()
	lis, err := net.Listen("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("fatal error:", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection\n")

	//常规启动http服务
	http.Serve(lis, nil)
}
