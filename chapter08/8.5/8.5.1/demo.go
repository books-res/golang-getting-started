package main

import "fmt"

func main() {
	var str = "uvwxyz"
	if len(str) >= 3 {
		goto L1
	} else {
		goto L2
	}

	// 第一个标签
L1:
	fmt.Println("字符串的长度符合要求")
	return

	// 第二个标签
L2:
	fmt.Println("字符串的长度不足3个字节")
}