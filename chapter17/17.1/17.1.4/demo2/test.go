package main

import (
	"fmt"
	"os"
)

func main() {
	// 以只读模式打开文件
	file, err := os.OpenFile("somefile", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 尝试写入
	data := []byte{21, 32, 100, 96, 14, 70, 52, 65, 10, 28}
	n, err := file.Write(data)
	if err != nil {
		fmt.Printf("写入错误：%v\n", err)
		file.Close() // 关闭文件
		return
	}
	fmt.Printf("已向文件写入%d个字节\n", n)
}
