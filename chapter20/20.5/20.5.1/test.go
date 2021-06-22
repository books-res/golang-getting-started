package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开.zip文件
	zreader, err := zip.OpenReader("testfile.zip")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读出里面的文件
	for _, f := range zreader.File {
		fmt.Printf("文件：%s\n", f.Name)
		// 打开文件流
		freader, err := f.Open()
		if err != nil {
			fmt.Printf("错误：%v\n", err)
			continue
		}
		fmt.Print("内容：")
		// 将文件内容复制到标准输出流
		io.Copy(os.Stdout, freader)
		// 关闭文件流
		freader.Close()
		fmt.Print("\n\n")
	}

	// 关闭.zip文件
	zreader.Close()
}
