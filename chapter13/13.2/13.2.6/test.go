package main

import "fmt"

func main() {
	var x = 35000
	fmt.Printf("十进制：%d\n", x)
	fmt.Printf("二进制：%#b\n", x)
	fmt.Printf("八进制：%#o\n", x)
	fmt.Printf("十六进制：%#x\n", x)
}
