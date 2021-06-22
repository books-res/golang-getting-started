package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer = bytes.NewBuffer(nil)
	// 压缩五个文件
	gzw := gzip.NewWriter(buffer)
	for n := 1; n <= 5; n++ {
		// 设置文件名
		gzw.Name = fmt.Sprintf("file-%02d.txt", n)
		// 写入文件内容
		var content = fmt.Sprintf("示例文本 -- %d\n", n)
		gzw.Write([]byte(content))
		// 写入完成
		gzw.Close()
		// 重置，准备写入下一个文件
		gzw.Reset(buffer)
	}

	// 解压
	gzr, err := gzip.NewReader(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 逐一读出文件
	for {
		// 此处调用是关键
		gzr.Multistream(false)
		fmt.Printf("文件：%s\n", gzr.Name)
		fmt.Print("文件内容：")
		// 读出文件内容，并写入标准输出流
		io.Copy(os.Stdout, gzr)
		fmt.Print("\n")
		// 重置，读取下一个文件
		err := gzr.Reset(buffer)
		if err == io.EOF {
			// 到达文档末尾，退出循环
			break
		}
	}
	// 关闭Reader对象
	gzr.Close()
}
