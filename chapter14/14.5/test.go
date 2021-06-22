package main

import (
	"fmt"
	"math"
)

func main() {
	// 开平方根
	n := float64(400)
	r := math.Sqrt(n)
	fmt.Printf("%.0f开平方根后：%.0f\n", n, r)
	// 开立方根
	n = 125
	r = math.Cbrt(n)
	fmt.Printf("%.0f开立方根后：%.0f\n", n, r)
}
