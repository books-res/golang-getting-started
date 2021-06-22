package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"../helper"
)

func main() {
	// 从命令行参数中本地侦听地址
	if len(os.Args) != 2 {
		fmt.Println("未提供服务器的侦听地址")
		return
	}
	var svAddr = os.Args[1]

	// 开始侦听
	listener, err := net.Listen("tcp", svAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("等待客户端连接……")
	// 接受连接，并接收文件
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		claddr := conn.RemoteAddr()
		fmt.Printf("客户端【%s】已连接\n", claddr)
		// 在新的协程上接收文件
		go RecFile(conn)
	}
}

func RecFile(c net.Conn) {
	defer c.Close()
	// 文件名
	mFilename := readFilename(c)
	//fmt.Printf("收到的文件：%#v\n", mFilename)
	// 校验码
	checkSum := readChecksum(c)
	// 保存文件
	readFile(mFilename, c)
	// 计算保存文件的校验码
	newFlchecksum, err := helper.CheckSum(mFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 比较校验码
	isChecked := compareChecksum(newFlchecksum, checkSum)
	if isChecked {
		fmt.Printf("文件%s校验成功\n", mFilename)
	} else {
		fmt.Printf("文件%s校验失败\n", mFilename)
	}
}

// 读取文件名
func readFilename(c net.Conn) string {
	data := make([]byte, 100)
	n, _ := c.Read(data)
	var tf = string(data[:n])
	// 当文件名的长度小于100字节
	// 剩余的内容会用0填充
	// 为了得到正确的文件名
	// 需要把字符串末尾的填充字符去掉
	return strings.TrimRight(tf, "\x00")
}

// 读取校验码
func readChecksum(c net.Conn) []byte {
	data := make([]byte, sha1.Size)
	n, _ := c.Read(data)
	return data[:n]
}

// 比较校验码
func compareChecksum(a, b []byte) bool {
	for i := 0; i < sha1.Size; i++ {
		if a[i] != b[i] {
			// 只要有一个字节不相等
			// 就表明校验码不一致
			return false
		}
	}
	return true
}

// 读取文件
func readFile(filename string, c net.Conn) {
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写入文件内容
	io.Copy(outFile, c)
	outFile.Close()
	fmt.Printf("文件%s已保存\n", filename)
}
