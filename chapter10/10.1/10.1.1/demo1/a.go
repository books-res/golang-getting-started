package main

import "fmt"

func main() {
	// 为元素分配的默认值

	var x [4]byte		// 0
	fmt.Println(x)

	var y [3]bool		// false
	fmt.Println(y)

	var z [3]string		// ""
	fmt.Println(z)
}