package main

import "fmt"

func main() {
	// 四个常量
	const Val1 int = 0
	const Val2 int = 1
	const Val3 string = "SPEED"
	const Val4 bool = false

	// 输出
	fmt.Printf("Val1: %v\n", Val1)
	fmt.Printf("Val2: %v\n", Val2)
	fmt.Printf("Val3: %v\n", Val3)
	fmt.Printf("Val4: %v", Val4)

	// 常量不允许修改
	//Val1 = 5
	//Val2 = 6

	// 可以自动推判类型
	const LockModeA = -1
	const LockModeB = 1
	const OrgChars = "ZXYkhDf6PngHeQR1LjdyU7a"

	// 输出
	fmt.Printf("\nLockModeA: %v\t%[1]T\n", LockModeA)
	fmt.Printf("LockModeB: %v\t%[1]T\n", LockModeB)
	fmt.Printf("OrgChars: %v\t%[1]T", OrgChars)
}