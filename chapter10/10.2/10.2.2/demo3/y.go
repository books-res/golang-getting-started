package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5}
	fmt.Printf("初始元素列表：%v\n", s)

	// 截取除最后一个元素外的所有元素
	s = s[0:len(s) - 1]
	fmt.Printf("删掉最后一个元素后：%v\n", s)
}