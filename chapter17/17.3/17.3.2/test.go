package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建新文件
	file, err := os.Create("Srcfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入内容
	var data = []byte{1, 2, 3, 4, 5}
	file.Write(data)
	// 关闭文件
	file.Close()

	// 创建符号链接
	err = os.Symlink("Srcfile", "myLink")
	if err != nil {
		fmt.Println(err)
		return
	}
}
