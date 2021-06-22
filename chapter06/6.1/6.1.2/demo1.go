package main

import (
	"fmt"
)

func main() {
	var st string = "zyx"
	fmt.Println(st)

	// 字符串中出现双引号时需要转义
	var sc = "My name is \"Tom\""
	fmt.Println(sc)
}