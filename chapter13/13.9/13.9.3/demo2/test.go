package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "www.example.org"
	var s2 = strings.Trim(s, "www.org")
	fmt.Printf("处理前：%s\n", s)
	fmt.Printf("处理后：%s\n", s2)
}
