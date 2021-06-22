package main

import (
	"fmt"

	"./demo"
)

func main() {
	fmt.Println("等待结果……")
	demo.Start()
	var x = <-demo.C
	fmt.Println("结果出来了")
	fmt.Printf("结果：%d\n", x)
}
