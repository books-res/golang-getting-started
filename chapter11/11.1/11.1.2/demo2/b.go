package main

import "fmt"

func main() {
	var xm = map[int32][]byte{ }
	xm[21] = []byte("c7g59rof71j5")
	xm[43] = []byte("xyxyxy")
	fmt.Printf("第一次赋值（key: 43）：%v\n", xm[43])
	// 相同的key会替换原有的值
	xm[43] = []byte("dkdkdk")
	fmt.Printf("第二次赋值（key: 43）：%v\n", xm[43])
}