package main

import (
	"fmt"
)

func main() {
	var a uint8 = 220
	var b uint8 = 89
	var c = a | b
	fmt.Printf("%08[1]b(%[1]d) | %08[2]b(%[2]d): %08[3]b(%[3]d)\n", a, b, c)

	// 声明变量并初始化
	var (
		h = 0b_001
		i = 0b_010
		j = 0b_100
	)
	// 组合运算
	fmt.Printf("\ni与h组合：%03b\n", i | h)
	fmt.Printf("i与j组合：%03b\n", i | j)
	fmt.Printf("h、i、j三者组合：%03b", h | i | j)
}