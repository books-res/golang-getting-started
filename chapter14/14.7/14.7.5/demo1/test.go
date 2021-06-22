package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var list = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("原次序：\n%v\n\n", list)
	// 用于交换元素顺序的函数
	var swap = func(i, j int) {
		list[i], list[j] = list[j], list[i]
	}
	// 开始“洗牌”
	var ln = len(list)
	// 第一轮
	rand.Shuffle(ln, swap)
	fmt.Printf("第一轮：\n%v\n\n", list)
	// 第二轮
	rand.Shuffle(ln, swap)
	fmt.Printf("第二轮：\n%v\n\n", list)
	// 第三轮
	rand.Shuffle(ln, swap)
	fmt.Printf("第三轮：\n%v\n\n", list)
}
