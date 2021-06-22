package main

import (
	"fmt"
)

func main() {
	// 第三个值被丢弃
	a,b,_ := "abc","lmn","opq"

	fmt.Printf("a: %v\nb: %v\n", a,b)

	// 调用函数，第三个返回被丢弃
	h,k,_ := comp(8)
	fmt.Print(h,k)
}
 
func comp(n int32) (int32, int32, bool) {
	var x1, x2 int32
	x1 = n * 2	// 乘以2
	x2 = n * n	// 二次方
	var b bool
	if n % 2 == 0 {
		b = true	// 偶数
	} else {
		b = false	// 奇数
	}
	return x1,x2,b
}
