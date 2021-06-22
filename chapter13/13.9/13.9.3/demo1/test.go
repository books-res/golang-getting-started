package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "上海自来水来自海上"
	var s2 = strings.Trim(s, "海水上来自")
	fmt.Printf("修剪前：%q\n", s)
	fmt.Printf("修剪后：%q\n", s2)
}
