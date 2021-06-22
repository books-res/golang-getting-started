package main

import "fmt"

func main() {
	var b uint32
	for b = 100; b < 500; b += 50 {
		if b == 300 {
			break		// 退出整个循环
		}
		fmt.Println(b)
	}
}