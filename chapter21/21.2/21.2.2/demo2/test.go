package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建通道实例
	var mych = make(chan string, 1)

	go func() {
		fmt.Println("开始执行新的协程")
		// 向通道发送数据
		mych <- "Hello"
		// 再发送一次就会阻塞
		//mych <- "World"
		fmt.Println("新协程执行完毕")
	}()

	// 暂停一下
	time.Sleep(time.Second * 2)
	fmt.Println("主协程即将退出")
}
