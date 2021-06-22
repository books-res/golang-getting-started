package main

import (
	"fmt"
)


func main() {
	// 声明变量
	var s string
	// 输出默认值
	// 默认值是 nil，输出后不可见
	//fmt.Printf("变量 s 的默认值：%s\n", s)

	// 修改变量的值
	s = "你好"
	// 输出变量的值
	//fmt.Printf("变量 s 的当前值：%v\n", s)

	// 多次赋值
	s = "早上好"
	s = "中午好"
	s = "下午好"
	s = "晚上好"
	// 变量 s 的当前值为“晚上好”
	fmt.Printf("变量 s 的最新值为：%v\n", s)

	// 声明变量 b ，并赋值（初始化）
	var b int16 = 680
	fmt.Printf("变量 b 的值为：%v\n", b)

	// 省略变量类型
	var c = 3.14159
	fmt.Printf("变量 c 的值为：%v，类型为%[1]T\n", c)

	var d = float32(3.14159)
	// 或者
	//var d float32 = 3.14159
	fmt.Printf("变量 d 的值为：%v，类型为：%[1]T\n", d)

	// 简练语法
	f := "xyz"
	g := 1.5e7
	fmt.Printf("变量 f 的值：%v，类型为：%[1]T\n", f)
	fmt.Printf("变量 g 的值为：%v，类型为：%[1]T\n", g)
}