package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = []byte("Some Data")
	// 编码
	// 确定编码后会产生多少个字节
	n := base64.StdEncoding.EncodedLen(len(data))
	encodeData := make([]byte, n)
	base64.StdEncoding.Encode(encodeData, data)
	// 输出
	fmt.Printf("原数据：%#x\n", data)
	fmt.Printf("编码后：%#x\n", encodeData)
	// 解码
	// 先确定解码后的字节个数
	n = base64.StdEncoding.DecodedLen(len(encodeData))
	decodeData := make([]byte, n)
	base64.StdEncoding.Decode(decodeData, encodeData)
	fmt.Printf("解码后：%#x\n", decodeData)
}
