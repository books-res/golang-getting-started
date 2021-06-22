package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	// 创建文件
	file, err := os.Create("encoded.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 准备编码
	var b64Writer = base64.NewEncoder(base64.StdEncoding, file)
	// 写入要编码的内容
	// 可分多次写入
	b64Writer.Write([]byte("春江潮水连海平\n"))
	b64Writer.Write([]byte("海上明月共潮生\n"))
	b64Writer.Write([]byte("滟滟随波千万里\n"))
	b64Writer.Write([]byte("何处春江无月明\n"))
	b64Writer.Write([]byte("江流宛转绕芳甸\n"))
	b64Writer.Write([]byte("月照花林皆似霰\n"))
	// 释放资源
	b64Writer.Close()
	file.Close()

	fmt.Print("编码完毕\n\n")

	// 打开文件
	file, err = os.Open("encoded.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 准备解码
	var decoder = base64.NewDecoder(base64.StdEncoding, file)
	// 解码
	fmt.Print("解码后：\n")
	io.Copy(os.Stdout, decoder)
	// 释放资源
	file.Close()
}
