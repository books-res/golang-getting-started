package demo

import "time"

// 此变量仅在包内访问
var innerch = make(chan int, 1)

// 此变量对外公开
var C <-chan int = innerch

// 此函数对外公开
func Start() {
	time.Sleep(time.Second * 2)
	innerch <- 10000
}
