package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 第一个参数，类型：int
	p1 := reflect.TypeOf(0)
	// 第二个参数，类型：bool
	p2 := reflect.TypeOf(true)
	// 第三个参数
	// 此参数个数可变，类型：[]string
	p3 := reflect.TypeOf([]string{})

	// 第一个返回值，类型：int
	rt1 := reflect.TypeOf(0)
	// 第二个返回值，类型：bool
	rt2 := reflect.TypeOf(true)

	// 构造输入/输出参数列表
	pin := []reflect.Type{p1, p2, p3}
	pout := []reflect.Type{rt1, rt2}

	// 创建函数类型
	t := reflect.FuncOf(pin, pout, true)
	fmt.Println("函数：", t)
}
