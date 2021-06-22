package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"../helper"
)

func main() {
	// 命令行参数应该有两个：
	// 1、服务器的连接地址
	// 2、要发送的文件路径
	// 第一个元素是可执行程序的路径
	// 参数从第二个元素开始处理
	if len(os.Args) != 3 {
		fmt.Println("命令行参数不完整")
		return
	}

	// 服务器地址
	var svhost = os.Args[1]
	// 待发送的文件
	var fileToSend = os.Args[2]

	// 连接服务器
	conn, err := net.Dial("tcp", svhost)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 发送文件名
	filename := helper.GetBaseName(fileToSend)
	data := make([]byte, 100)
	copy(data, []byte(filename))
	conn.Write(data)

	// 发送校验码
	cksum, err := helper.CheckSum(fileToSend)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write(cksum)

	// 发送文件
	inFile, err := os.Open(fileToSend)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(conn, inFile)
	// 关闭文件
	inFile.Close()
	fmt.Println("文件发送完毕")
	// 关闭通信连接
	conn.Close()
}
