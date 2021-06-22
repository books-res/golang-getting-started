package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var ints = rand.Perm(5)
	for _, v := range ints {
		fmt.Println(v)
	}
}
