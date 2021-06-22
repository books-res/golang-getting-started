package main

import "fmt"

type char = int32

func main() {
	var x char = 'H'
	var y int32 = 72
	// char与int32是同一类型

	fmt.Printf("x和y变量相等？%t", x == y)
}