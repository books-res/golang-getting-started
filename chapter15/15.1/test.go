package main

import (
	"fmt"
	"sort"
)

func main() {
	// 整数列表的排序
	fmt.Println("------- 整数列表 -------")
	var intNums = []int{780, -5, 20, 375, 196, 86, -455, 67}
	fmt.Printf("排序前：%v\n", intNums)
	sort.Ints(intNums)
	fmt.Printf("排序后：%v\n", intNums)

	fmt.Println()
	// 浮点数列表的排序
	fmt.Println("------- 浮点数列表 -------")
	var floatNums = []float64{0.01472, -2.881652, 4.13, 0.0021, 19.7789, 5.99}
	fmt.Printf("排序前：%v\n", floatNums)
	sort.Float64s(floatNums)
	fmt.Printf("排序后：%v\n", floatNums)

	fmt.Println()
	// 字符串列表的排序
	fmt.Println("------- 字符串列表 -------")
	var strList = []string{"zero", "bus", "six", "just", "as", "flash"}
	fmt.Printf("排序前：%v\n", strList)
	sort.Strings(strList)
	fmt.Printf("排序后：%v\n", strList)
}
