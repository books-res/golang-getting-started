package main

import "fmt"

func main() {
	var x interface{} = float64(0.00123)
	// 执行断言
	y := x.(float64)
	fmt.Printf("变量x的真实类型为：%T\n", y)
	fmt.Printf("它的值为：%v\n", y)

	// 断言失败
	//z := x.(int8)
	//fmt.Printf("变量z的值：%v\n", z)

	// 断言失败，但不会发生错误
	z, ok := x.(int8)
	if ok {
		fmt.Printf("断言成功，z的值为：%v\n", z)
	} else {
		fmt.Println("断言失败")
	}
}