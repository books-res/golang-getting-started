package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 输入参数
	var (
		p1   = reflect.TypeOf(float32(0.01))
		p2   = reflect.TypeOf(float32(0.01))
		pins = []reflect.Type{p1, p2}
	)

	// 返回值
	var ret = reflect.TypeOf("")
	var pouts = []reflect.Type{ret}

	// 创建函数类型
	myFunc := reflect.FuncOf(pins, pouts, false)

	// 输出一下函数类型
	fmt.Printf("刚创建的函数类型：\n%v\n\n", myFunc)

	// 创建函数体
	funcInst := reflect.MakeFunc(myFunc, func(in []reflect.Value) []reflect.Value {
		// 取出两个参数的值
		f1 := in[0].Float()
		f2 := in[1].Float()
		// 将两个值相减
		fr := f1 - f2
		// 将结果转换为字符串
		str := fmt.Sprintf("%.2f", fr)
		// 封装返回值
		retv := reflect.ValueOf(str)
		return []reflect.Value{retv}
	})

	// 调用函数
	var input1, input2 float32 = 0.9, 0.7
	cr := funcInst.Call([]reflect.Value{
		reflect.ValueOf(input1),
		reflect.ValueOf(input2),
	})
	// 输出结果
	var resval = cr[0].String()
	fmt.Printf("输入参数：%.2f，%.2f\n", input1, input2)
	fmt.Printf("函数调用结果：%s\n", resval)
}
