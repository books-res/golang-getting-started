package main

import "fmt"

func main() {
	if a := 3; a > 0 {
		fmt.Print("a的值：", a)		// 第一条Print语句
	}

	// 错误：此范围无法访问变量a
	//fmt.Print("a的值：", a)		// 第二条Print语句
}