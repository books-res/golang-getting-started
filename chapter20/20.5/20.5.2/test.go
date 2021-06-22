package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// 在内存中暂存数据的缓冲区
	var buffer = new(bytes.Buffer)

	// 构建新的Zip文档
	zipw := zip.NewWriter(buffer)
	// 添加四个文件
	file1, _ := zipw.Create("a.dat")
	file1.Write([]byte("red red red red red"))
	file2, _ := zipw.Create("b.dat")
	file2.Write([]byte("tick tick tick"))
	file3, _ := zipw.Create("c.dat")
	file3.Write([]byte("core-core-core"))
	file4, _ := zipw.Create("d.dat")
	file4.Write([]byte("test*test*test*test*test"))
	// 关闭Writer
	zipw.Close()

	// 读出Zip文档中的文件
	baseReader := bytes.NewReader(buffer.Bytes())
	zipr, err := zip.NewReader(baseReader, baseReader.Size())
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取文件列表
	for _, f := range zipr.File {
		fmt.Printf("文件：%s\n", f.Name)
		fmt.Printf("大小：%d\n", f.UncompressedSize64)
		fmt.Printf("压缩后大小：%d\n", f.CompressedSize64)
		fmt.Print("文件内容：")
		// 读取文件内容
		fr, err := f.Open()
		if err != nil {
			continue
		}
		io.Copy(os.Stdout, fr)
		// 关闭文件
		fr.Close()
		fmt.Print("\n\n")
	}
}
