package main

import "fmt"

func main() {
	var a string
	var b int
	fmt.Scanf("%s%d", &a, &b)

	/*
		有效输入：
		abc 123
		abc1 23
		12 3
		无效输入：
		123abc
		abc123
		123bc
	*/
	fmt.Printf("a= %s, b= %d\n", a, b)
}
