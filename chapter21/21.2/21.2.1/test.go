/*
	注意：本示例只演示通道的基本用法
	不要直接运行代码，因为无缓冲的
	通道在同一协程中使用会发生死锁
*/

package main

func main() {
	//var c = make(chan string)
	var c = make(chan string, 0)
	// 将数据发送到通道
	c <- "hello"
	// 接收通道中的数据
	<-c
	// 关闭通道
	close(c)
}
