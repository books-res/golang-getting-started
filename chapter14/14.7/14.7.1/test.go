package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("------ 32位浮点数 ------")
	for i := 0; i < 5; i++ {
		var n = rand.Float32()
		fmt.Println(n)
	}

	fmt.Println("\n------ 64位浮点数 ------")
	for i := 0; i < 5; i++ {
		var n = rand.Float64()
		fmt.Println(n)
	}
}
