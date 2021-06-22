package main

import "fmt"

func main() {
	// 常规用法
	/*
	var a = 10
	var b = 20
	var c = 30
	*/

	// 组合赋值
	var a,b,c = 10,20,30
	// 输出到屏幕上
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 简约语法
	x,y,z := "test",5,0.002
	// 输出
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// 错误：“=”左右两边的表达式个数不等
	//q1,q2,q3,q4 := -400,"bit",0.2

	// 错误：提供了过多的值
	//v1,v2 := -1,-2,-3

	r1,r2,r3 := test()
	fmt.Println("\n函数返回的三个值：")
	fmt.Printf("r1: %v\nr2: %v\nr3: %v", r1,r2,r3)
}

// 示例函数
func test() (string, string, int) {
	return "abc", "xyz", 10000
}