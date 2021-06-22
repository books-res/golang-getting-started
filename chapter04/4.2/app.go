package main

import (
	"fmt"
	"calculator"
)

func main() {
	var num1, num2 = 15, 20
	// 依次调用 Add、Sub、Mult 函数
	var (
		r1 = calculator.Add(num1, num2)
		r2 = calculator.Sub(num1, num2)
		r3 = calculator.Mult(num1, num2)
	)
	// 打印结果
	fmt.Printf("%d + %d = %d\n", num1, num2, r1)
	fmt.Printf("%d - %d = %d\n", num1, num2, r2)
	fmt.Printf("%d * %d = %d\n", num1, num2, r3)
}