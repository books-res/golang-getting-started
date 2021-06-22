package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x1 float64 = -23.00
		x2 float64 = 18.02
		x3 float64 = 15
		x4 float64 = -3000
		x5 float64 = 0.56
	)

	// 计算绝对值
	r1 := math.Abs(x1)
	r2 := math.Abs(x2)
	r3 := math.Abs(x3)
	r4 := math.Abs(x4)
	r5 := math.Abs(x5)

	// 输出结果
	fmt.Printf("%.2f的绝对值：%.2f\n", x1, r1)
	fmt.Printf("%.2f的绝对值：%.2f\n", x2, r2)
	fmt.Printf("%.2f的绝对值：%.2f\n", x3, r3)
	fmt.Printf("%.2f的绝对值：%.2f\n", x4, r4)
	fmt.Printf("%.2f的绝对值：%.2f\n", x5, r5)
}
