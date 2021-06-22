package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "0xb20d8a"
	var n, err = strconv.ParseInt(str, 0, 32)
	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%q -> %v\n", str, int32(n))

	str = "111"
	n2, err := strconv.ParseUint(str, 2, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%q -> %v\n", str, uint8(n2))

	str = "0.20205"
	n3, err := strconv.ParseFloat(str, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%q -> %.6f\n", str, float32(n3))
}
