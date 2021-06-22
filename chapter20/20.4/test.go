package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func main() {
	// 构造字典
	var dict = []byte(`{
	"id": ,
	"name": "",
	"age": ,
	"city": ""
}`)

	// 待压缩文本
	var srcText = `{
	"id": 32001,
	"name": "Tommy",
	"age": 27,
	"city": "Beijing"
}`
	var srcData = []byte(srcText)
	fmt.Printf("原数据长度：%d\n", len(srcData))

	var bf bytes.Buffer
	// 压缩
	zw, _ := flate.NewWriterDict(&bf, 9, dict)
	zw.Write(srcData)
	zw.Close()
	fmt.Printf("压缩后数据长度：%d\n", bf.Len())

	// 修改字典
	// for i := range dict {
	// 	dict[i] = '&'
	// }

	// 解压缩
	zr := flate.NewReaderDict(&bf, dict)
	fmt.Print("解压后：")
	io.Copy(os.Stdout, zr)
	zr.Close()
}
