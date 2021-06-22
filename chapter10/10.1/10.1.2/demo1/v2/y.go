package main

import "fmt"

func main() {
	var x = [5]rune{'a', 'e', 'i', 'o', 'u'}
	// 获取数组的长度
	n := len(x)
	// 最后一个元素的索引为n-1，倒数第二个的索引为n-2
	last1 := x[n-2]
	last2 := x[n-1]
	fmt.Printf("最后两个元素：%c、%c\n", last1, last2)
}