package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "data.txt"
	// 打开文件
	var file, err = os.Open(fileName)
	if err != nil {
		fmt.Printf("错误：%v\n", err)
		return
	}
	// 当代码执行完当前上下文后自动释放资源
	defer file.Close()
	// 从文件读入文本
	fmt.Println("从文件中读到的内容如下：")
	err = nil
	// 存储读入的一行文本
	var line string
	// 统行读入的行数
	var lineNo uint
	for {
		_, err = fmt.Fscanln(file, &line)
		// 行数递增
		lineNo++
		if err == io.EOF {
			// 已到达文件末尾，跳出循环
			break
		}
		fmt.Printf("%-4d %s\n", lineNo, line)
	}
}
