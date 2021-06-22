package main

import (
	"fmt"
	"sort"
)

func main() {
	// 初始化切片对象
	var nums = []int{17, 81, 3, 437, 69, 13, 97}
	fmt.Printf("切片：%v\n", nums)

	// 用于递增排序的函数
	var increaseFun = func(i, j int) bool {
		return nums[i] < nums[j]
	}
	// 用于递减排序的函数
	var decreaseFun = func(i, j int) bool {
		return !(nums[i] < nums[j])
	}

	sort.Slice(nums, increaseFun)
	fmt.Printf("递增：%v\n", nums)

	sort.Slice(nums, decreaseFun)
	fmt.Printf("递减：%v\n", nums)
}
