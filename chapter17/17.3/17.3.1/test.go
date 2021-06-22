package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// 创建F1文件
	myfile, err := os.Create("F1")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入文件内容
	myfile.WriteString("This is a file.\n")
	// 关闭文件
	myfile.Close()

	// 创建硬链接F2
	err = os.Link("F1", "F2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建硬链接F3
	err = os.Link("F2", "F3")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取F3的内容
	myfile, _ = os.Open("F3")
	fmt.Println("F3的内容：")
	bf := new(bytes.Buffer)
	io.Copy(bf, myfile)
	fmt.Printf("%s\n", string(bf.Bytes()))

	// 通过F2修改文件
	myfile, _ = os.OpenFile("F2", os.O_WRONLY|os.O_APPEND, 0644)
	myfile.WriteString("This is a text file.\n")
	myfile.Close()

	// 通过F1来读取内容
	myfile, _ = os.Open("F1")
	bf.Reset() // 清空缓冲区中的数据
	fmt.Println("F1的内容：")
	io.Copy(bf, myfile)
	fmt.Printf("%s\n", string(bf.Bytes()))
}
