package main

import (
	"fmt"
	"math"
)

func main() {
	// 坐标点
	var x, y float64 = 15.28, 54.72
	// 计算夹角
	rad := math.Atan2(y, x)
	// 转换为角度
	d := rad * 180 / math.Pi
	fmt.Printf("点 (%.2f, %.2f) 与x轴的夹角：%.2f\n", x, y, d)
}
