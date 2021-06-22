package main

import "fmt"

func main() {
	var n uint8 = 27
	var r = ^n
	fmt.Printf("原值：%[1]d(%08[1]b)\n", n)
	fmt.Printf("取反：%[1]d(%08[1]b)\n\n", r)

	var g uint16 = 150
	var q = ^g
	fmt.Printf("原值：%[1]d(%016[1]b)\n", g)
	fmt.Printf("取反：%[1]d(%016[1]b)\n", q)
}