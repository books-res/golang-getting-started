package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const fileName = "demo.txt"
	// 数据来源为文件流
	var reader, err = os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入目标为标准输出流
	var writer = os.Stdout

	fmt.Println("------ 从文件中复制的内容 ------")
	// 复制文件内容
	_, err = io.Copy(writer, reader)
	if err != nil {
		fmt.Println(err)
	}
}
