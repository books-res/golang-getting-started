package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	// 输入文件，待压缩
	inFile, _ := os.Open("music.mp3")
	// 输出文件，已压缩
	outFile, _ := os.Create("music.gz")

	gzw := gzip.NewWriter(outFile)
	// 设置在压缩包内显示的文件名
	gzw.Name = "music.mp3"
	// 压缩并写入数据
	io.Copy(gzw, inFile)
	// 关闭Writer对象并把数据写入文件
	gzw.Close()
	// 关闭文件
	inFile.Close()
	outFile.Close()
}
