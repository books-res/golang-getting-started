package main

import "fmt"

func main() {
	test1(16)							// 可变参数个数：0
	test1(24, "abcd")					// 可变参数个数：1
	test1(3, "Jack", "Dick", "Tom")		// 可变参数个数：3

	test3()
	test3(0.001)
	test3(20.22, 16.033, 0.7, 1.55)
	// 也可以这样调用
	nums := []float32 { 0.002, 13.5, 0.17 }
	test3(nums...)
}

func test1(a uint8, b ...string) {
	fmt.Printf("参数a：%d\n", a)
	fmt.Printf("参数b：%v\n", b)
}

// 下面的函数定义有错误
/*
func test2(p1 string, p2 ...bool, p3 float32) {
	// ……
}
*/

func test3(args ...float32) {
	n := len(args)
	fmt.Printf("\n\n可变参数个数：%d\n", n)
	// 列出可变参数中的元素
	if n > 0 {
		fmt.Println("参数内容：")
		for _,val := range args {
			fmt.Printf("%f  ", val)
		}
	}
}