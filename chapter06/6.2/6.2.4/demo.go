package main

import "fmt"

func main() {
	// 128位
	var ca complex128 = 50 + 2i
	var cb complex128 = -0.05 - 3i
	fmt.Println(ca)
	fmt.Println(cb)

	// 64位
	var cc complex64 = 1.0001 + 0.0005i
	var cd complex64 = -300 + 12i
	fmt.Println(cc)
	fmt.Println(cd)

	// 先写虚部，后写实部
	var t complex64 = 0.3i + 15
	fmt.Println(t)

	// 大写的 i 是错误的
	//var a complex128 = 100 + 6I
	//fmt.Println(a)

	// 分隔符
	var t2 complex128 = 1_333 + 0.2_635i
	fmt.Println(t2)

	// 分隔符的错误用法
	//var t3 complex64 = 82 + 0.412_i
}