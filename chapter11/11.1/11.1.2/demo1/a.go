package main

import "fmt"

func main() {
	var mt = make(map[int32]byte)
	// 添加/设置元素
	mt[1] = 25
	mt[2] = 26
	mt[3] = 27
	// 读取元素
	val1 := mt[1]
	val2 := mt[2]
	val3 := mt[3]
	fmt.Println(val1, val2, val3)
}