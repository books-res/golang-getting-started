package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 自定义字符映射表
	var encodeStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@#"
	// 创建encoding
	custEncoding := base64.NewEncoding(encodeStr)

	// 待编码的数据
	testData := "一二三四五六七"
	// 编码
	var ecStr = custEncoding.EncodeToString([]byte(testData))
	fmt.Printf("原数据：%s\n", testData)
	fmt.Printf("编码后：%s\n", ecStr)

	// 解码
	var decodeData, _ = custEncoding.DecodeString(ecStr)
	fmt.Printf("解码后：%s\n", decodeData)
}
