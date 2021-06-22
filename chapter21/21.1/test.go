package main

import (
	"fmt"
)

var (
	A = make(chan int)
	B = make(chan int)
	C = make(chan int)
	D = make(chan int)
)

func test1() {
	fmt.Print("春")
	A <- 1
}

func test2() {
	<-A
	fmt.Print("夏")
	B <- 1
}

func test3() {
	<-B
	fmt.Print("秋")
	C <- 1
}

func test4() {
	<-C
	fmt.Print("冬")
	D <- 1
}

func main() {
	// 启动四个新的协程
	go test1()
	go test2()
	go test3()
	go test4()

	// 暂停一下
	//time.Sleep(time.Second)
	<-D
}
