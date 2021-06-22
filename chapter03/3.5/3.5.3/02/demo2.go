package main

import "fmt"

func main() {
	var a, b, c int8
	var res int8

	a = 7
	res = ^a
	fmt.Printf("原值：%[1]d(%08[1]b)\n", a)
	fmt.Printf("取反：%[1]d(%08[1]b)\n\n", res)

	b = -6
	res = ^b
	fmt.Printf("原值：%[1]d(%08[1]b)\n", b)
	fmt.Printf("取反：%[1]d(%08[1]b)\n\n", res)

	c = -13
	res = ^c
	fmt.Printf("原值：%[1]d(%08[1]b)\n", c)
	fmt.Printf("取反：%[1]d(%08[1]b)", res)
}