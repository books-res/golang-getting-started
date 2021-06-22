package main

import (
	"fmt"
)

func main() {
	// 声明两个变量
	var x, y int

	// 第一次
	x = 121
	y = 16
	fmt.Printf("%d %% %d = %d\n", x, y, x % y)

	// 第二次
	x = -90
	y = -18
	fmt.Printf("%d %% %d = %d\n", x, y, x % y)

	// 第三次
	x = 175
	y = -29
	fmt.Printf("%d %% %d = %d", x, y, x % y)
}
