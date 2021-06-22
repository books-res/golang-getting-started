package main

import (
	"fmt"
	"math/big"
)

func main() {
	var strfloat = "0.001234567890234567893456789"
	var bigFloat = new(big.Float)
	// 设置精度
	bigFloat.SetPrec(50)
	// 设置浮点数值
	bigFloat.SetString(strfloat)
	// 打印到屏幕上
	fmt.Printf("%.50f\n", bigFloat)
}
