package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 创建新的文件
	var file, err = os.Create("test.data")
	if err != nil {
		fmt.Printf("错误：%v\n", err)
		return
	}
	fmt.Println("成功创建test.data文件")

	// 随机产生字节序列
	rand.Seed(time.Now().Unix())
	buffer := make([]byte, 100)
	rand.Read(buffer)
	fmt.Printf("已生成随机字节序列：\n%x\n", buffer)

	// 写入文件
	file.Write(buffer)
	// 关闭文件引用
	file.Close()
	fmt.Println("文件写入完毕\n")

	//-----------------------------------
	// 打开文件
	fileToRead, err := os.Open("test.data")
	if err != nil {
		fmt.Printf("错误：%v\n", err)
		return
	}
	// 读取内容
	fmt.Println("正在读取文件……")
	fmt.Print("文件内容：\n")
	readBuf := make([]byte, 16)
	for {
		n, err := fileToRead.Read(readBuf)
		if err == io.EOF {
			break
		}
		fmt.Printf("%x", readBuf[:n])
	}
}
