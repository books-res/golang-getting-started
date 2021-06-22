package main

import "fmt"

func main() {
	var x = [5]rune{'a', 'e', 'i', 'o', 'u'}
	
	// 倒数第二个元素，索引为3
	last1 := x[3]
	// 最后一个元素，索引为4
	last2 := x[4]
	fmt.Printf("最后两个元素：%c、%c\n", last1, last2)
}