package main

import "fmt"

func main() {
	fmt.Println("----- 不带换行符 -----")
	fmt.Print("第一项")
	fmt.Print("第二项")
	fmt.Print("第三项")
	fmt.Print("第四项")

	fmt.Println("\n\n----- 带换行符 -----")
	fmt.Println("第一项")
	fmt.Println("第二项")
	fmt.Println("第三项")
	fmt.Println("第四项")
}
