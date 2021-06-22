package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer = new(bytes.Buffer)

	// 打包Tar文档
	tarw := tar.NewWriter(buffer)
	// 第一个文件
	content := []byte("第一个文件")
	header := &tar.Header{
		Name: "p01.txt",
		Size: int64(len(content)), //内容长度
	}
	// 写入文件头
	if err := tarw.WriteHeader(header); err == nil {
		// 写入内容
		tarw.Write(content)
	}
	// 第二个文件
	content = []byte("第二个文件")
	header.Name = "p02.txt"
	header.Size = int64(len(content))
	if err := tarw.WriteHeader(header); err == nil {
		tarw.Write(content)
	}
	// 第三个文件
	content = []byte("第三个文件")
	header.Name = "p03.txt"
	header.Size = int64(len(content))
	if err := tarw.WriteHeader(header); err == nil {
		tarw.Write(content)
	}
	// 第四个文件
	content = []byte("第四个文件")
	header.Name = "p04.txt"
	header.Size = int64(len(content))
	if err := tarw.WriteHeader(header); err == nil {
		tarw.Write(content)
	}
	// 关闭Writer对象
	tarw.Close()

	// 解包Tar文档
	tarr := tar.NewReader(buffer)
	// 循环读取文件列表
	for {
		hd, err := tarr.Next()
		if err == io.EOF {
			// 已到了列表末尾，跳出循环
			break
		}
		// 打印文件名
		fmt.Printf("文件：%s\n", hd.Name)
		// 读取文件内容
		fmt.Print("内容：")
		io.Copy(os.Stdout, tarr)
		fmt.Print("\n\n")
	}
}
