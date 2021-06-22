package main

import (
	"fmt"
)

func main() {
	var x1 rune = 'f' // 小写英文字母
	var x2 = 'G'      // 大写英文字母
	var x3 = '好'      // 中文字符
	var x4 rune = ':' // 标点符号
	var x5 rune = '@' // 特殊符号
	var x6 rune = '7' // 数字字符
	// 单引号需要转义
	var x7 = '\''

	/*
		注意：rune类型的变量直接输出的是整数值
	*/
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(x5)
	fmt.Println(x6)
	fmt.Println(x7)
}
