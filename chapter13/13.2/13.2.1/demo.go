package main

import "fmt"

func main() {
	var num int32 = 6037029
	fmt.Println("int32原数值：", num)
	fmt.Printf("二进制：%b\n", num)
	fmt.Printf("八进制：%o\n", num)
	fmt.Printf("十进制：%d\n", num)
	fmt.Printf("十六进制（小写）：%x\n", num)
	fmt.Printf("十六进制（大写）：%X\n", num)

	fmt.Printf("我叫%s，今年%d岁\n", "小明", 21)
}
