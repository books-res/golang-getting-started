package main

import "fmt"

// 声明三个变量
var (
	k = 0.0001
	j = 0.0021
	m int16 = 5530
)

// 声明三个常量
const (
	XldFirst = "F"
	XldSecond = "G"
	XldThird = "H"
)

func main() {
	fmt.Println("变量：")
	fmt.Println("k:", k)
	fmt.Println("h:", j)
	fmt.Println("m:", m)
	fmt.Println("\n常量：")
	fmt.Println("XldFirst:", XldFirst)
	fmt.Println("XldSecond:", XldSecond)
	fmt.Println("XldThird:", XldThird)	
}