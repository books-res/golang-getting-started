package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	// 待压缩的文本
	var text = `
{
	"id": 372144,
	"pno": "FC-B10-23-1",
	"in_date": "2018-11-25",
	"desc": "Kismil",
	"color": "Black"	
}`
	// 原数据的字节序列
	var srcBs = []byte(text)

	// 使用gzip算法压缩数据
	var buffer = bytes.NewBuffer(nil)
	gzw := gzip.NewWriter(buffer)
	// 写入数据
	gzw.Write(srcBs)
	// 调用Close方法让压缩后的数据写入基础数据流中
	// 调用后不会关闭基础数据流
	gzw.Close()

	// 解压缩数据
	gzr, err := gzip.NewReader(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("解压缩后的内容：")
	// 读取解压缩的内容
	io.Copy(os.Stdout, gzr)
	// 关闭Reader对象
	// 与其关联的基础数据流不会关闭
	gzr.Close()
}
