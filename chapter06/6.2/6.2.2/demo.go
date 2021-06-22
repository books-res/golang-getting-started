package main

import (
	"fmt"
)

func main() {
	var (
		a uint = 8000				// 十进制
		b int = 0b11101				// 二进制
		c int32 = 0o102532			// 八进制
		d = 0x6afc20e			// 十六进制（小写）
		e = 0XC33EE4F			// 十六进制（大写）
	)
	fmt.Printf("%d\t%d\t%d\t%d\t%d\n", a,b,c,d,e)

	// 使用分隔符
	var (
		h uint64 = 257_6888_102_453
		i int64 = -800_22520_16_3
		j uint64 = 0b_10001_1111_00111110
		k uint64 = 0x_ff_e6_41_c32a
	)
	fmt.Printf("%d\t%d\t%d\t%d\n", h,i,j,k)

	// 以下代码会出错
 	//var (
	//	w = _7600_240
	//	v = 0xe5e3fc20_
	//)
 }