package main

import "fmt"

func main() {
	// 接收返回值
	var result = add(12, 7)
	fmt.Println(result)
	// 丢弃返回值
	add(9, 320)
	// 以下调用，12传给参数y，7传给参数x
	//var result = add(7, 12)

	// 调用无返回值的函数
	sendTo("102.33.0.51", 5353)
}

func add(x, y int16) int16 {
	return x + y
}

func sendTo(host string, port int32) {
	// ……
}