package main

import (
	"fmt"
	"math"
)

func main() {
	var d float64 = 30.0
	// 转换为弧度角
	var r = d * math.Pi / 180
	// 计算正弦值
	var s = math.Sin(r)
	// 计算余弦值
	var c = math.Cos(r)
	fmt.Printf("角度：%.2f\n", d)
	fmt.Printf("正弦值：%.2f\n余弦值：%.2f\n", s, c)
}
