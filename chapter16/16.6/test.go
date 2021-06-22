package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 原始数据
	var dataSrc = strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	fmt.Println("原数据：")
	io.Copy(os.Stdout, dataSrc)
	// 截取 off = 12，n = 7
	rd := io.NewSectionReader(dataSrc, 12, 7)
	fmt.Println("\n\n第一次截取：")
	io.Copy(os.Stdout, rd)
	// 截取 off = 30，n = 13
	rd = io.NewSectionReader(dataSrc, 30, 13)
	fmt.Println("\n\n第二次截取：")
	io.Copy(os.Stdout, rd)
}
