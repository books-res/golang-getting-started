package main

import (
	"fmt"
)

func main() {
	var d0 rune = '月'
	var d1 rune = '亮'
	var d2 rune = 'g'
	var d3 rune = 'Z'

	// 输出
	// %v 格式控制符只输出整数值，而非字符
	//fmt.Printf("%v %v %v %v", d0, d1, d2, d3)
	// 输出原字符
	fmt.Printf("%c  %c  %c  %c", d0, d1, d2, d3)

	// 以下代码会发生错误
	//var d4 rune = 'abc'	
}