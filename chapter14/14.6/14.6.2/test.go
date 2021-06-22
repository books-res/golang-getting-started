package main

import (
	"fmt"
	"math/big"
)

func main() {
	var c = new(big.Int)
	// 求30的阶乘
	c.MulRange(1, 30)
	fmt.Printf("30! = %d\n", c)
	// 求50的阶乘
	c.MulRange(1, 50)
	fmt.Printf("50! = %d\n", c)
}
