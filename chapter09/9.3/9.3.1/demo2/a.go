package main

import "fmt"

type test interface {
	sendMessage(head, body string) int
}

func main() {
	var ix test
	fmt.Printf("接口类型变量的默认值：%v", ix)
}