package main

import "fmt"

func checkType(val interface{}) {
	switch x := val.(type) {
	case uint8:
		fmt.Printf("无符号8位整数：%d\n", x)
	case float32:
		fmt.Printf("32位浮点数：%f\n", x)
	case string:
		fmt.Printf("字符串：%s\n", x)
	case func():
		fmt.Println("这是个函数，无参数，无返回值")
	case func(string) int64:
		fmt.Println("这是个函数，输入参数为字符串类型，返回值为64位有符号整数")
	default:
		fmt.Println("其他类型")
	}
}

func main() {
	// 变量列表
	var (
		v1 = "jeep"
		v2 = func(){ }
		v3 = uint8(24)
		v4 = struct { K int } { 350 }
	)
	// 调用函数
	checkType(v1)
	checkType(v2)
	checkType(v3)
	checkType(v4)
}