package main

import (
	"fmt"
	"math"
)

func main() {
	var d float64 = 45.0
	// 转换为弧度角
	r := math.Pi / 180 * d
	// 求正弦、余弦值
	var s, c = math.Sincos(r)
	fmt.Printf("角度：%.1f\n", d)
	fmt.Printf("正弦值：%.1f\n", s)
	fmt.Printf("余弦值：%.1f\n", c)
}
