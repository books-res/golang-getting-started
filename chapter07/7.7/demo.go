package main

import "fmt"

func main() {
	var fnx = func(a, b float32) (float32, float32) {
		return a*a, b*b
	}
	work(fnx)

	// 定义三个匿名函数
	var pf1 = func(i int) {
		fmt.Printf("二进制：%#b\n", i)
	}
	var pf2 = func(i int) {
		fmt.Printf("十进制：%d\n", i)
	}
	var pf3 = func(i int) {
		fmt.Printf("十六进制：%#x\n", i)
	}

	// 调用函数
	printElements(pf1, 5, 8, 12)
	printElements(pf2, 5, 8, 12)
	printElements(pf3, 5, 8, 12)
}

// 也可以这样写：func printElements(fn func(x int), items ...int)
func printElements(fn func(int), items ...int) {
	if len(items) == 0 || fn == nil {
		return
	}
	for _,n := range items {
		fn(n)
	}
}