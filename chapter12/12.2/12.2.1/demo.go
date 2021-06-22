package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 32位无符号整数
	var a uint32
	fmt.Print("变量a：")
	checkType(a)

	// 指针类型（*uint32）
	var pa = &a
	fmt.Print("变量pa：")
	checkType(pa)

	// 函数类型
	var b = func() { }
	fmt.Print("变量b：")
	checkType(b)

	// 匿名的结构体
	var c struct{ X,Y int32; Z rune }
	fmt.Print("变量c：")
	checkType(c)

	// 布尔类型
	var d bool
	fmt.Print("变量d：")
	checkType(d)

	// 通道类型（chan）
	var e chan <- int 
	fmt.Print("变量e：")
	checkType(e)

	// 映射类型
	var f = map[uint]string {1:"abc", 2:"opq", 3:"jkl"}
	fmt.Print("变量f：")
	checkType(f)

	// 数组类型
	var g = [2]string{ "time", "code" }
	fmt.Print("变量g：")
	checkType(g)

	// 切片类型
	var h = g[:]
	fmt.Print("变量h：")
	checkType(h)
}

func checkType(o interface{}) {
	var tp = reflect.TypeOf(o)
	// 分析类型
	switch tp.Kind() {
	case reflect.Bool:
		fmt.Println("布尔类型")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("有符号整数")
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Println("无符号整数")
	case reflect.Float32, reflect.Float64:
		fmt.Println("浮点数")
	case reflect.String:
		fmt.Println("字符串")
	case reflect.Struct:
		fmt.Println("结构体")
	case reflect.Interface:
		fmt.Println("接口")
	case reflect.Array:
		fmt.Println("数组")
	case reflect.Slice:
		fmt.Println("切片")
	case reflect.Map:
		fmt.Println("映射")
	case reflect.Complex64, reflect.Complex128:
		fmt.Println("复数")
	case reflect.Ptr:
		fmt.Println("指针")
	case reflect.Func:
		fmt.Println("函数")
	case reflect.Chan:
		fmt.Println("通道")
	default:
		fmt.Println("其他类型")
	}
}