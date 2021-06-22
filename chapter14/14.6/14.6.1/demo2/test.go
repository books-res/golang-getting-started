package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 两个整数值的字符串形式
	var (
		num1 = "8000006554217002143024"
		num2 = "6676225236988563328930"
	)
	// 实例化Int对象，使用指针类型（*Int）
	var bigInt1 = new(big.Int)
	var bigInt2 = new(big.Int)
	// 将两个字符串表示的值设置到Int对象中
	bigInt1.SetString(num1, 10)
	bigInt2.SetString(num2, 10)
	// 进行加法运算
	var res1 = new(big.Int)
	res1.Add(bigInt1, bigInt2)
	fmt.Printf("%d + %d = %d\n", bigInt1, bigInt2, res1)
	// 进行减法运算
	var res2 = new(big.Int)
	res2.Sub(bigInt1, bigInt2)
	fmt.Printf("%d - %d = %d\n", bigInt1, bigInt2, res2)
}
