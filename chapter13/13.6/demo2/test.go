package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "dog dog dog"
	str2 := strings.ReplaceAll(str, "g", "t")
	fmt.Printf("原字符串：%s\n", str)
	fmt.Printf("替换之后：%s\n", str2)
}
