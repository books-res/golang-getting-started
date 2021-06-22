package main

import (
	"fmt"
	"os"
)

func main() {
	// 移动文件
	err := os.Rename("fd/first/abc.txt", "fd/second/abc.txt")
	if err != nil {
		fmt.Println("错误：", err)
	}
}
