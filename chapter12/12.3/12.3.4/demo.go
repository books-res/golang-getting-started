package main

import (
	"fmt"
	"reflect"
)

// 定义函数
func add(m, n int32) int32 {
	return m + n
}

func sub(p, q int32) int32 {
	return p - q
}


func main() {
	var a1, a2 int32 = 15, 17
	var r1 = callFunc(add, a1, a2)
	fmt.Printf("add(%v, %v) => %v\n", a1, a2, r1)

	a1, a2 = 30, 12
	var r2 = callFunc(sub, a1, a2)
	fmt.Printf("sub(%v, %v) => %v\n", a1, a2, r2)
}

func callFunc(fun interface{}, args ...interface{}) []interface{} {
	fv := reflect.ValueOf(fun)
	if fv.Kind() != reflect.Func {
		fmt.Println("被调用的不是函数")
		return []interface{}{ }
	}
	// 传入的参数个数
	inlen := len(args)
	// 被调用函数的参数个数
	funptLen := fv.Type().NumIn()
	// 检查传入参数的个数是否正确
	if inlen != funptLen {
		fmt.Println("传入的参数个数不正确")
		return []interface{}{ }
	}
	// 检查参数类型是否正确
	for i := 0; i < inlen; i++ {
		ti := reflect.TypeOf(args[i])
		tfi := fv.Type().In(i)
		if ti.Kind() != tfi.Kind() {
			fmt.Println("参数类型不正确")
			return []interface{}{ }
		}
	}
	// 调用目标函数
	// 提取参数值
	var prts = make([]reflect.Value, inlen)
	for i := 0; i < inlen; i++ {
		prts[i] = reflect.ValueOf(args[i])
	}
	var res = fv.Call(prts)
	// 提取返回值
	outlen := len(res)
	if outlen == 0 {
		return []interface{}{ }
	}
	var outs = make([]interface{}, outlen)
	for i := 0; i < outlen; i++ {
		outs[i] = res[i].Interface()
	}
	return outs
}