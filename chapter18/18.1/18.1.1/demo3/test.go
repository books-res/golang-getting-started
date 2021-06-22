package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var test = "<1, 2, 3, 4, 5>"
	// Base64编码
	var ec = base64.URLEncoding.EncodeToString([]byte(test))
	// 演示URL
	fmt.Printf("原数据：%s\n", test)
	fmt.Printf("编码后：%s\n", ec)
	fmt.Printf("URL样例：http://someone.com/nav?pages=%s\n", ec)
}
