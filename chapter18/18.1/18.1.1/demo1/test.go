package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 原数据
	var testData = "Some Data"
	// 编码
	var encodeStr = base64.StdEncoding.EncodeToString([]byte(testData))
	// 输出结果
	fmt.Printf("原数据：%s\n", testData)
	fmt.Printf("Base64编码后：%s\n", encodeStr)
	// 解码
	var decodeData, err = base64.StdEncoding.DecodeString(encodeStr)
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	fmt.Printf("解码后：%s\n", string(decodeData))
}
