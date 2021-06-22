package main

import "fmt"

func main() {
	var key = 'E'
	switch key {
	case 'A':
		fmt.Println("选项一")
	case 'B':
		fmt.Println("选项二")
	case 'C':
		fmt.Println("选项三")
	default:
		fmt.Println("未知选项")
	}
}