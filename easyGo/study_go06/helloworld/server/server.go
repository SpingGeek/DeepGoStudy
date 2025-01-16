// Package server @Author Spring
// @Date 2025/1/3 21:34:00
// @Desc
package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改 reply 的值
	*reply = "hello. " + request
	return nil
}

func main() {
	// 1.实例化一个 server
	listener, _ := net.Listen("tcp", ":1234")
	// 2.注册处理逻辑 handler
	err := rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		return
	}
	// 3.启动服务
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)

}
