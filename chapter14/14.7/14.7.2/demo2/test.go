package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 7; i++ {
		n := rand.Intn(100)
		fmt.Println(n)
	}
}
