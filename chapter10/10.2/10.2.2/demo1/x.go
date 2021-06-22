package main

import (
	"fmt"
)

func main() {
	var sf = []float32{0.001, 0.0007}
	printSliceInfo(sf)
	// 添加一个元素
	sf = append(sf, 0.0014)
	printSliceInfo(sf)
	// 添加两个元素
	sf = append(sf, 0.0008, 0.1205)
	printSliceInfo(sf)
	// 添加三个元素
	sf = append(sf, 0.0275, 1.302, 5.0071)
	printSliceInfo(sf)
}

func printSliceInfo(s []float32) {
	fmt.Printf("长度：%d，容量：%d，元素列表：%v\n", len(s), cap(s), s)
}