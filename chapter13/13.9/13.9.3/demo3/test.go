package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "id_562043"
	var s2 = strings.TrimRight(s, "0123456789")
	fmt.Printf("处理前：%s\n", s)
	fmt.Printf("处理后：%s\n", s2)
}
