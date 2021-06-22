package main

import (
	"fmt"
)

func main() {
	// 声明变量
	var (
		m int8 = 12
		n int8 = 6
	)
	// 按位与
	result := m & n
	fmt.Printf("%08b & %08b: %08b (%[3]d)\n", m, n, result)
}