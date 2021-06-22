package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 种子相同，每次运行所生成的随机序列相同
	for x := 0; x < 6; x++ {
		fmt.Println(rand.Int())
	}
}
