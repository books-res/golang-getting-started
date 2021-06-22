package main

import "fmt"

func main() {
	var s = make([]string, 0, 10)
	fmt.Printf("初始化后，长度：%d，容量：%d\n", len(s), cap(s))
	// 添加50个元素
	for i := 1; i <= 50; i++ {
		str := fmt.Sprintf("Item %d", i)
		s = append(s, str)
	}
	fmt.Printf("添加50个元素后，长度：%d，容量：%d\n", len(s), cap(s))
}