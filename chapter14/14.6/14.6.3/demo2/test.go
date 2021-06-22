package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 两个原始数值，类型均为64位浮点数
	var (
		a float64 = 1550.797220660354896132160489549963189232654585896489465456159657
		b float64 = 0.0016200166894105953690156
	)
	var (
		bigFa  = new(big.Float)
		bigFb  = new(big.Float)
		bigRes = new(big.Float) // 存放计算结果
	)
	// 设置数值精度
	bigFa.SetPrec(100)
	bigFb.SetPrec(100)
	bigRes.SetPrec(100)
	// 设置原始数值
	bigFa.SetFloat64(a)
	bigFb.SetFloat64(b)

	// 加
	bigRes.Add(bigFa, bigFb)
	fmt.Printf("%.15f + %.15f = %.15f\n", bigFa, bigFb, bigRes)

	// 减
	bigRes.Sub(bigFa, bigFb)
	fmt.Printf("%.15f - %.15f = %.15f\n", bigFa, bigFb, bigRes)

	// 乘
	bigRes.Mul(bigFa, bigFb)
	fmt.Printf("%.15f * %.15f = %.15f\n", bigFa, bigFb, bigRes)

	// 除
	bigRes.Quo(bigFa, bigFb)
	fmt.Printf("%.15f / %.15f = %.15f\n", bigFa, bigFb, bigRes)
}
