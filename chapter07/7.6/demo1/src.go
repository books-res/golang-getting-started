package main

import "fmt"

func main() {
	var myfun = func (x, y int) int {
		return x * x + y * y
	}

	// 调用
	res := myfun(2, 4)
	fmt.Printf("x² + y² = %d", res)
}