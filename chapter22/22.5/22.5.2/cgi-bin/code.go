package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 写入消息头
	fmt.Print("Content-Type: text/html; charset=utf-8\r\n")
	fmt.Print("Server-Ver: 1.0\r\n")
	var bts = make([]byte, 15)
	rand.Read(bts)
	fmt.Printf("Access-Token: %x\r\n", bts)
	// 消息头与正文之间有空行分隔
	fmt.Print("\r\n")
	// 以下是正文
	fmt.Print("<h3>这是一个简单的 CGI 程序</h3>")
}
