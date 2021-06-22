package main

import "net/http"

func main() {
	// 创那Server实例
	demoServer := new(http.Server)
	// 设置侦听地址
	demoServer.Addr = ":8080"
	// 开始侦听连接
	demoServer.ListenAndServe()
}
