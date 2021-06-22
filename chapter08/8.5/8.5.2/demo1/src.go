package main

import "fmt"

func main() {
	// 实例化三维数组
	var myarr = [3][3][3]uint{
		{
			{ 1, 2, 3 },
			{ 4, 5, 6 },
			{ 7, 8, 9 },
		},
		{
			{ 10, 11, 12},
			{ 13, 14, 15 },
			{ 16, 17, 18 },
		},
		{
			{ 19, 20, 21 },
			{ 22, 23, 24 },
			{ 25, 26, 27 },
		},
	}

	// 使用嵌套循环来枚举数组中的元素
	for _, a := range myarr {				// 第一层
		for _, b := range a {				// 第二层
			for _, c := range b {			// 第三层
				if c % 3 == 0 {
					continue
				}
				fmt.Printf("%d  ", c)
			}
		}
	}
}