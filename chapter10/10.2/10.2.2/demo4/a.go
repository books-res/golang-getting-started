package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("初始元素列表：%v\n", s)
	// 删去前两个元素
	s = s[2:]
	fmt.Printf("删除前两个元素后：%v\n", s)
}