package main

import (
	"fmt"
)

func main() {
	// 声明两个变量，用于存放操作数
	var m, n int
	// 让用户自行输入操作数
	fmt.Print("请输入第一个操作数：")
	fmt.Scan(&m)
	fmt.Print("请输入第二个操作数：")
	fmt.Scan(&n)
	// 进行四则运算
	r1 := m + n
	r2 := m - n
	r3 := m * n
	r4 := m / n
	// 输出计算结果
	fmt.Println("\n计算结果：")
	fmt.Printf("%d + %d = %d\n", m, n, r1)
	fmt.Printf("%d - %d = %d\n", m, n, r2) 
	fmt.Printf("%d * %d = %d\n", m, n, r3)
	fmt.Printf("%d / %d = %d", m, n, r4)
}