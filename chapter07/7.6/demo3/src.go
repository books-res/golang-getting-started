package main

import "fmt"

func main() {
	// 创建新的通道对象
	var ch = make(chan byte)
	go func () {
		fmt.Println("新协程")
		// 向通道对象发送数据
		ch <- 1
	}()

	// 从通道对象接收数据
	<- ch
	fmt.Println("主协程")
}