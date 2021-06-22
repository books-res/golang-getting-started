package main

import (
	"fmt"
	"sort"
)

// 定义新类型
type Int32Nums []int32

func (x Int32Nums) Less(i, j int) bool {
	// 求余
	m1 := x[i] % 8
	m2 := x[j] % 8
	return m1 < m2
}
func (x Int32Nums) Len() int {
	return len(x)
}
func (x Int32Nums) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	var dx = Int32Nums{27, 102, 58, 47, 85}
	fmt.Printf("原列表：%v\n", dx)
	sort.Sort(dx)
	fmt.Printf("排序后：%v\n", dx)
}
