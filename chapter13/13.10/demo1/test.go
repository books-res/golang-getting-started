package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "ABC"
	var s2 = strings.Repeat(str, 3)
	fmt.Printf("原字符串：%s\n", str)
	fmt.Printf("重复三次：%s\n", s2)
}
