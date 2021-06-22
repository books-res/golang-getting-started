package main

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func main() {
	// 输出文件名
	var testfile = "data.zip"
	outFile, err := os.Create(testfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	zw := zip.NewWriter(outFile)
	// 注册压缩算法
	zw.RegisterCompressor(zip.Deflate, func(w io.Writer) (io.WriteCloser, error) {
		// 最高压缩比
		return flate.NewWriter(w, flate.BestCompression)
	})
	// 添加三个文件
	tmp, _ := zw.Create("part1.txt")
	tmp.Write([]byte("test---------data---"))
	tmp, _ = zw.Create("part2.txt")
	tmp.Write([]byte("ab cd ab cd"))
	tmp, _ = zw.Create("part3.txt")
	tmp.Write([]byte("...content..."))
	// 关闭Writer
	zw.Close()

	// 关闭文件
	outFile.Close()
}
