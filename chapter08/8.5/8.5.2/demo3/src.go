package main

import "fmt"

func main() {
	// 创建切片实例
	var arr = [][]rune{
		{'a','k','f'},
		{'x','p'},
		{'z','y','s'},
	}

	fmt.Println("--------- 未使用标签 ---------")
	// 使用break语句退出循环，无代码标签
	for i, s := range arr {
		for j, c := range s {
			if c == 'f' {
				break
			}
			fmt.Printf("[%d][%d]: %c\n", i, j, c)
		}
	}

	fmt.Println("---------- 使用标签 ----------")
	// 使用break语句退出循环，最外层for循环有代码标签
	LT:for i, s := range arr {
		for j, c := range s {
			if c == 'f' {
				break LT
			}
			fmt.Printf("[%d][%d]: %c\n", i, j, c)
		}
	}
}