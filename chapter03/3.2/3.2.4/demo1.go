package main

import (
	"fmt"
)

func main() {
	var a = 100
	a++
	fmt.Printf("a++ -> %d\n", a)

	var b = 100
	b--
	fmt.Printf("b-- -> %d", b)
}