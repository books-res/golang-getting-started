package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建新文件
	file, err := os.Create("file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入内容
	file.WriteString("test")
	file.Close()

	// 重命名文件
	err = os.Rename("file1.txt", "file2.txt")
	if err != nil {
		fmt.Println(err)
	}
}
