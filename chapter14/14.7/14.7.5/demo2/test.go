package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := len(list) // 元素个数
	fmt.Printf("原顺序：%v\n", list)
	// 用于交换元素顺序的函数
	swap := func(i, j int) {
		if i != 0 && i != (n-1) && j != 0 && j != (n-1) {
			list[i], list[j] = list[j], list[i]
		}
	}
	// 开始洗牌
	fmt.Println("------ 洗牌 ------")
	// 设置种子值
	rand.Seed(time.Now().Unix())
	for x := 1; x <= 5; x++ {
		fmt.Printf("第%d次：", x)
		rand.Shuffle(n, swap)
		fmt.Printf("%v\n", list)
	}
}
