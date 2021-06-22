package main

import "fmt"

func main() {
	var (
		width = 7 // 宽度
		prec  = 2 // 精度
	)
	var x float32 = -0.085216

	// 第一次输出
	fmt.Printf("%[1]*.[2]*f\n", width, prec, x)

	// 修改宽度
	width = 12
	// 第二次输出
	fmt.Printf("%[1]*.[2]*f\n", width, prec, x)

	// 修改宽度和精度
	width = 20
	prec = 5
	// 第三次输出
	fmt.Printf("%[1]*.[2]*f\n", width, prec, x)
}
