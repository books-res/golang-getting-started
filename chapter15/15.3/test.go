package main

import (
	"fmt"
	"sort"
)

func main() {
	var strs = []string{"ststst", "ffff", "dxd", "egfegfegfegf", "kp", "u"}
	fmt.Printf("原来的字符串列表：\n%#v\n\n", strs)

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	fmt.Printf("按字符串长度递增排序后：\n%#v\n", strs)
}
