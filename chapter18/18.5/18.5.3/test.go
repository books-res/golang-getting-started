package main

import (
	"encoding/pem"
	"os"
)

func main() {
	// 消息正文
	var msg = "演示数据"
	// 消息头
	var headers = map[string]string{
		"ver":    "1.0",
		"sender": "Jack",
		"copyto": "Tom",
	}

	// 编码
	block := pem.Block{
		Type:    "EMSG",
		Bytes:   []byte(msg),
		Headers: headers,
	}
	pem.Encode(os.Stdout, &block)
}
