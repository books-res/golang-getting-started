package main

import "fmt"

func main() {
	// 为所有元素初始化
	var f = [5]uint32{ 18, 75, 42, 3, 105}
	// 仅初始化第一个元素
	//var f = [5]uint32{ 18, }
	// 或者
	//var f = [5]uint32{ 18 }
	fmt.Println(f)

	// 声明后通过索引初始化元素
	var r [4]float32
	r[0] = 1.112		// 第一个元素
	r[1] = 0.000054		// 第二个元素
	r[2] = 370.303		// 第三个元素
	r[3] = -16.75		// 第四个元素
}