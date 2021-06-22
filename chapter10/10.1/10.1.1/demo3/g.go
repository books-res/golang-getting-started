package main

import "fmt"

func main() {
	// 错误：元素类型不一致
	//a := [2]uint{ 1.7, 33 }

	// 空接口类型可兼容任何类型
	s := [3]interface{}{"abc", 887, 'H'}
	fmt.Println(s)
}