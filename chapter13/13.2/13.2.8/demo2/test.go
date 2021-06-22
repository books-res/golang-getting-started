package main

import "fmt"

func main() {
	var v = 1234.8765432

	fmt.Printf("宽度为20，精度为3：%20.3f\n", v)
	fmt.Printf("宽度为6，左对齐，精度为2：%-6.2f\n", v)
}
