package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("------ 产生int类型的随机整数 ------")
	for i := 0; i < 6; i++ {
		in := rand.Int()
		fmt.Println(in)
	}

	fmt.Println()
	fmt.Println("------ 产生int32类型的随机整数 ------")
	for i := 0; i < 6; i++ {
		in := rand.Int31()
		fmt.Println(in)
	}
}
