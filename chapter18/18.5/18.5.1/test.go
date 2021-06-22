package main

import (
	"encoding/pem"
	"fmt"
)

func main() {
	var content = "小明发来一张贺卡"

	// 编码
	var block pem.Block
	// 消息类型为“DEMO”
	block.Type = "DEMO"
	// 消息内容
	block.Bytes = []byte(content)
	// 生成编码后的数据
	var encodeData = pem.EncodeToMemory(&block)
	// 输出编码后的PEM消息
	fmt.Printf("PEM数据：\n%s\n\n", encodeData)

	// 解码
	var decblock, _ = pem.Decode(encodeData)
	// 输出
	fmt.Print("解码后的消息：\n")
	fmt.Printf("消息类型：%s\n", decblock.Type)
	fmt.Printf("消息正文：%s\n", decblock.Bytes)
}
