package main

import "fmt"

func main() {
	// 实例化一个数组对象
	src := [4]uint32{10, 20, 30, 40}
	// 从数组产生两个切片实例
	s1 := src[0:2]
	s2 := src[1:4]

	fmt.Println("------ 修改数组前 ------")
	fmt.Printf("数组：%v\n", src)
	fmt.Printf("切片1：%v\n", s1)
	fmt.Printf("切片2：%v\n", s2)

	// 修改数组中的元素
	src[0] = 100
	src[2] = 300
	fmt.Println("\n------ 修改数组后 ------")
	fmt.Printf("数组：%v\n", src)
	fmt.Printf("切片1：%v\n", s1)
	fmt.Printf("切片2：%v\n", s2)

	// 修改切片中的元素
	s1[1] = 700
	s2[2] = 900
	fmt.Println("\n------ 修改切片后 ------")
	fmt.Printf("数组：%v\n", src)
	fmt.Printf("切片1：%v\n", s1)
	fmt.Printf("切片2：%v\n", s2)
}