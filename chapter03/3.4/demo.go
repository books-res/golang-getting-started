package main

import (
	"strings"
	"fmt"
)

func main() {
	var b bool
	// 逻辑“与”
	b = 5 > 7 && 3 < 6
	fmt.Printf("5 > 7 && 3 < 6: %t\n", b)
	// 多个条件的逻辑“与”
	var (
		x = 18
		y = 5
		z = x % y
	)
	b = x <= 20 && y == 5 && z > 5
	fmt.Printf("x <= 20 && y == 5 && z > 5: %t\n", b)
	
	// 逻辑“或”
	b = strings.Contains("abcd", "ab") || 50 < 90 || 10 > 15
	fmt.Printf("abcd contains ab || 50 < 90 || 10 > 15: %t\n", b)

	// 逻辑“非”
	b = !(1 == 1)
	fmt.Printf("!(1 == 1): %t\n", b)
}