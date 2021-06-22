package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = "一天又一天"
	var b = strings.Replace(a, "天", "年", 2)
	fmt.Printf("原字符串：%s\n", a)
	fmt.Printf("替换之后：%s\n", b)
}
