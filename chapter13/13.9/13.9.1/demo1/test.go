package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "ftp://am:321@192.168.1.13"
	var str2 = strings.TrimPrefix(str, "ftp://")
	fmt.Printf("修剪前：%s\n修剪后：%s\n", str, str2)
}
