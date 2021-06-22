package main

import "fmt"

func main() {
	var r1 = 3 > 5
	fmt.Printf("3大于5吗？%t\n", r1)
	var r2 = 100%13 != 0
	fmt.Printf("100 ÷ 13 有余数吗？%t\n", r2)
}
