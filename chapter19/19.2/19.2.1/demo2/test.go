package main

import (
	"flag"
	"fmt"
)

func main() {
	// 变量列表
	var (
		// 第一个数值
		num1 float64
		// 第二个数值
		num2 float64
		// 运算方式
		opt string
	)

	// 绑定命令行参数
	flag.Float64Var(&num1, "x", 0.00, "第一个操作数")
	flag.Float64Var(&num2, "y", 0.00, "第二个操作数")
	flag.StringVar(&opt, "o", "+", "运算方式，仅支持四则运算。有效值为+、-、*、/")

	// 分析命令行参数
	flag.Parse()

	var result float64
	switch opt {
	case "+": //加法
		result = num1 + num2
	case "-": //减法
		result = num1 - num2
	case "*": //乘法
		result = num1 * num2
	case "/": //除法
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("错误：除数不能为0")
		}
	default:
		result = 0.00
		fmt.Println("输入的运算方式无效")
	}

	// 输出结果
	fmt.Printf("计算结果：%f\n", result)
}
