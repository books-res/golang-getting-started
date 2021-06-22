package main

import "fmt"

func main() {
	var number = 200
	switch number {
	case 200:
		// <- fallthrough语句不能出现在这里
		fmt.Println("分支一")
		fallthrough
	case 400:
		fmt.Println("分支二")
	case 600:
		fmt.Println("分支三")
	}
}