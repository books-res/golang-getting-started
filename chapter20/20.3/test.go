package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func main() {
	// 待压缩的字符串
	var testStr = "black, black, black, black, black"
	var srcData = []byte(testStr)
	fmt.Printf("原数据长度：%d\n", len(srcData))

	// 压缩
	var buffer = new(bytes.Buffer)
	fwt, err := flate.NewWriter(buffer, 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fwt.Write(srcData)
	// 调用此方法让数据写入基础流
	fwt.Close()

	fmt.Printf("压缩后数据长度：%d\n", buffer.Len())

	// 解压缩
	zrd := flate.NewReader(buffer)
	fmt.Print("解压后：")
	io.Copy(os.Stdout, zrd)
	zrd.Close()
}
