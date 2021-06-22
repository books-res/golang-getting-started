package main

import "fmt"

func example() {
	var q = 1

	for q < 10 {
		fmt.Printf("q的当前值：%d\n", q)
		q ++
	}
}

func example2() {
	// 这个循环永远不会执行
	for false {
		fmt.Println("循环条件为：false")
	}
}

func example3() {
	// 这个循环也是永远不会执行
	for 3 > 6 {
		fmt.Println("循环条件：3 > 6")
	}
}