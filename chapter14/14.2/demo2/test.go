package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a int32 = -900
		b int32 = -200
	)
	// 注意类型转换
	mx := int32(math.Max(float64(a), float64(b)))
	mn := int32(math.Min(float64(a), float64(b)))

	fmt.Printf("数值：%d、%d\n", a, b)
	fmt.Printf("最大值：%d\n", mx)
	fmt.Printf("最小值：%d\n", mn)
}
