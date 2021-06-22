package main

import "fmt"

func main() {
	// 创建通道实例
	var mych = make(chan int)
	// 在新的协程上运行
	go func() {
		select {
		case mych <- 1:
		case mych <- 2:
		case mych <- 3:
		case mych <- 4:
		case mych <- 5:
		}
	}()
	// 接收通道中的数据
	n := <-mych
	fmt.Printf("随机整数：%d\n", n)
}
