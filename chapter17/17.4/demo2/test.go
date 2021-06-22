package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var bts = []byte{0xa6, 0x57, 0xc5, 0xf3, 0xe9, 0x12}
	// 将内容写入文件
	ioutil.WriteFile("CKData", bts, 0666)
	// 从文件中读取内容
	var cnt, _ = ioutil.ReadFile("CKData")
	fmt.Printf("从文件中读到的内容：%#x\n", cnt)
}
