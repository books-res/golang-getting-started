package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建一个子目录
	err := os.Mkdir("data", 0700)
	if os.IsNotExist(err) {
		return
	}

	// 改变工作目录
	err = os.Chdir("data")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用工作目录
	file, err := os.Create("list")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入测试内容
	file.WriteString("demo")
	// 关闭文件
	file.Close()
}
