package main

import (
	"net"
	"net/http"
)

func main() {
	httpsvr := new(http.Server)
	// 创建Listener
	listener, err := net.Listen("tcp4", "127.0.0.1:88")
	if err != nil {
		return
	}
	// 调用Serve方法
	httpsvr.Serve(listener)
}
