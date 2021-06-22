package main

import "fmt"

func main() {
	var num1, num2 int // 参与运算的整数
	var op string      // 表示运算法则
	// 读入第一个整数
	fmt.Print("请输入第一个整数：")
	fmt.Scanln(&num1)
	// 读入第二个整数
	fmt.Print("请输入第二个整数：")
	fmt.Scanln(&num2)
	// 读入运算法则
	fmt.Print("请输入运算法则：")
	fmt.Scanln(&op)

	var result int
	// 分析并完成计算
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		// 如果输入的运算法则无效
		// 则默认进行加法运算
		op = "+"
		result = num1 + num2
	}

	// 输出结果
	fmt.Printf("\n计算结果：%d %s %d = %d\n", num1, op, num2, result)
}
