package main

import (
	"encoding/pem"
	"fmt"
)

func main() {
	var pemData = `
-----BEGIN MY NAME-----
546L5aSn5bGx
-----END MY NAME-----
Hello, Jim`

	var msgblock, other = pem.Decode([]byte(pemData))
	// 输出
	fmt.Printf("消息类型：%s\n", msgblock.Type)
	fmt.Printf("消息正文：%s\n", msgblock.Bytes)
	fmt.Printf("其他内容：%s\n", other)
}
