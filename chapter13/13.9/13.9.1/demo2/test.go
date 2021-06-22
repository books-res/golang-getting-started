package main

import (
	"fmt"
	"strings"
)

func main() {
	files := []string{
		"dx.xct",
		"sc_u.xct",
		"op_v6.xct",
		"etacl.xct",
	}

	fmt.Println("文件列表：")
	for _, f := range files {
		fmt.Println(f)
	}

	// 去除扩展名
	var prcdFiles = make([]string, 0)
	for _, f := range files {
		nf := strings.TrimSuffix(f, ".xct")
		prcdFiles = append(prcdFiles, nf)
	}

	fmt.Println("\n去掉扩展名后：")
	for _, f := range prcdFiles {
		fmt.Println(f)
	}
}
