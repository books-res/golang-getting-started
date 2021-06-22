package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "    山顶千门次第开    "
	str2 := strings.TrimSpace(str)
	fmt.Printf("原字符串：%#v\n", str)
	fmt.Printf("去掉首尾的空格：%#v\n", str2)
}
