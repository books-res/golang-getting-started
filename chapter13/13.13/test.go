package main

import (
	"fmt"
	"strings"
)

func main() {
	var bd = new(strings.Builder)
	// 开始构建字符串
	bd.WriteString("你好")
	bd.WriteRune('\n')
	bd.WriteString("小明")
	// 输出字符串
	fmt.Printf("刚创建的字符串：\n%s\n", bd)
}
