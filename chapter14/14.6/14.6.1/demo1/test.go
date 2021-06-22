package main

import "fmt"

func main() {
	// 错误
	var (
		a uint64 = 8000006554217002143024
		b uint64 = 6676225236988563328930
	)
	c1 := a + b
	c2 := a - b
	fmt.Printf("%d + %d = %d\n", a, b, c1)
	fmt.Printf("%d - %d = %d\n", a, b, c2)
}
