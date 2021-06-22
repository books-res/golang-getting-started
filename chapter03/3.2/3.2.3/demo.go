package main

import (
	"fmt"
	"math"
)

func main() {
	result := math.Pow(5, 3)
	fmt.Printf("5的3次方：%d\n", int(result))
	result = math.Pow(8, 2)
	fmt.Printf("8的2次方：%d\n", int(result))
	// 以10为底数
	result = math.Pow10(4)
	fmt.Printf("10的4次方：%d\n", int(result))
	// 开平方根
	result = math.Pow(81, 0.5)
	fmt.Printf("81开平方根：%d", int(result))
}