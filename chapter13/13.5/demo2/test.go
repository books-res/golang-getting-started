package main

import (
	"fmt"
	"strings"
)

func main() {
	var strs = []string{
		"SP",
		"LP",
		"CD",
		"VHS",
		"LD",
		"VCD",
		"DVD",
	}

	var out = strings.Join(strs, "#")
	fmt.Printf("连接后的字符串：%s\n", out)

	var cts = []string{"abcd", "efg", "hijk"}
	out = strings.Join(cts, "")
	fmt.Printf("空白分隔符：%s\n", out)
}
