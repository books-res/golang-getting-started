package main

import (
	"fmt"
	"reflect"
)

// 测试函数
func demoFunc(a string, b uint16, c float64) int32 {
	return 0
}

func main() {
	tp := reflect.TypeOf(demoFunc)
	
	// 参数信息
	// 获取参数个数
	var pmnum = tp.NumIn()
	// 枚举输入参数列表
	fmt.Println("----- 函数的输入参数 -----")
	for x := 0; x < pmnum; x++ {
		var p = tp.In(x)
		fmt.Printf("%d: %s\n", x, p.Name())
	}
	
	// 返回值信息
	// 获取返回值个数
	var rtnum = tp.NumOut()
	// 枚举返回值列表
	fmt.Println("\n----- 函数的返回值 -----")
	for x := 0; x < rtnum; x++ {
		r := tp.Out(x)
		fmt.Printf("%d: %s\n", x, r.Name())
	}
}