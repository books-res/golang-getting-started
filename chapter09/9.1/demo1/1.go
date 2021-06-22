package main

import "fmt"

// 从现有类型定义新类型
type name string

func main() {
	var a string = "abcde"
	var b name = "abcde"
	fmt.Printf("变量a的类型：%T\n", a)
	fmt.Printf("变量b的类型：%T\n", b)

	// 不能进行相等比较
	//a == b
}