package main

import "fmt"

func main() {
	var a = [5]int32{2, 4, 6, 8, 10}
	s := a[2:4]
	fmt.Println(s)
	s2 := a[2:5]
	fmt.Println(s2)
	// 截取数组中的所有元素
	s3 := a[:]
	fmt.Println(s3)
}