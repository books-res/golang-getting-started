package main

import "fmt"

func main() {
	arr1 := [4]float32{0.11, 0.23, 5.001, 12.63}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	fmt.Println()

	arr2 := [3]string{"zh-CN", "en-US", "zh-TW"}
	for index,value := range arr2 {
		fmt.Printf("索引：%d，值：%s\n", index, value)
	}
}