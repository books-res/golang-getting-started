package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		num1 float64 = 8515.33
		num2 float64 = 124.09
	)

	fmt.Printf("数值：%.2f、%.2f\n", num1, num2)
	fmt.Printf("较大的值：%.2f\n", math.Max(num1, num2))
	fmt.Printf("较小的值：%.2f\n", math.Min(num1, num2))
}
