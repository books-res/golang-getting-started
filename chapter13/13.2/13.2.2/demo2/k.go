package main

import "fmt"

func main() {
	var (
		a float32 = 1.234567e-3
		b float32 = 1.654321234e+7
	)
	fmt.Printf("%g\n%g\n", a, b)
}
