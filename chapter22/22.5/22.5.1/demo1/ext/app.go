package main

import (
	"fmt"
)

func main() {
	// 从标准输入流中读取两个值
	var a, b float32
	fmt.Scanln(&a, &b)
	// 计算结果
	var res = a * b
	// 将结果写入标准输出流
	fmt.Printf("计算结果：%.4f", res)
}
