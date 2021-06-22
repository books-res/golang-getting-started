package main

import (
	"fmt"
	"math"
)

func main() {
	// 计算3的4次方
	var bs, exp float64 = 3, 4
	result := math.Pow(bs, exp)
	fmt.Printf("%.0f的%.0f次方：%.0f\n", bs, exp, result)

	// 计算2的6次方
	bs, exp = 2, 6
	result = math.Pow(bs, exp)
	fmt.Printf("%.0f的%.0f次方：%.0f\n", bs, exp, result)

	// 计算20的-2次方
	bs, exp = 20, -2
	result = math.Pow(bs, exp)
	fmt.Printf("%.0f的%.0f次方：%.4f\n", bs, exp, result)
}
