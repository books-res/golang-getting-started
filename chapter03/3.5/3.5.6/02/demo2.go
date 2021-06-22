package main

import "fmt"

func main() {
	var (
		f1 uint8 = 0b_0000_0001    //香蕉
		f2 uint8 = 0b_0000_0010    //葡萄
		f3 uint8 = 0b_0000_0100    //芒果
		f4 uint8 = 0b_0000_1000    //秋梨
	)

	// 四种水果组合
	var all = f1 | f2 | f3 | f4
	fmt.Printf("四种水果组合：%04b\n", all)

	// 去掉部分水果
	fmt.Printf("去掉香蕉：%04b\n", all &^ f1)
	fmt.Printf("去掉芒果：%04b\n", all &^ f3)
	fmt.Printf("去掉芒果和秋梨：%04b", all &^ (f3 | f4))
}