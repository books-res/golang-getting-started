package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	myfile, err := os.Create("testdata")
	// 检查是否发生错误
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	// 写入内容
	var content = `第一行文本
第二行文本
第三行文本`
	myfile.WriteString(content)
	// 关闭文件
	myfile.Close()

	//------------------------------------
	myfile, err = os.Open("testdata")
	// 检查一下是否有错误
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	// 读取内容
	var data = new(bytes.Buffer)
	io.Copy(data, myfile)
	// 关闭文件
	myfile.Close()
	// 将读到的内容打印到屏幕上
	fmt.Printf("从文件中读到的内容：\n%s\n", data.Bytes())
}
