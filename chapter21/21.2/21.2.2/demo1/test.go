package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建通道实例
	var mych = make(chan uint)

	// 启动新的协程
	go func() {
		fmt.Println("开始执行新协程")
		// 向通道发送数据
		mych <- 350
		fmt.Println("新协程执行完毕")
	}()

	// 从通道中接收数据
	<-mych
	// 暂停一下
	time.Sleep(time.Second)
	fmt.Println("主协程即将退出")
}
