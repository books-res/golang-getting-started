package main

import "fmt"

func main() {
	var (
		n1 = uint(15)
		n2 = false
		n3 = complex64(27 - 3i)
		n4 = "hot"
		n5 = int16(-7)
		n6 = uint64(936600300)
		n7 = float32(-0.93)
		n8 = struct {
			f1 uint
			f2 byte
		}{500, 20} // 匿名结构体
		n9  = []byte("c4857da2")
		n10 = int32(88)
	)

	// 输出变量值的类型
	fmt.Printf("n1的数据类型：%T\n", n1)
	fmt.Printf("n2的数据类型：%T\n", n2)
	fmt.Printf("n3的数据类型：%T\n", n3)
	fmt.Printf("n4的数据类型：%T\n", n4)
	fmt.Printf("n5的数据类型：%T\n", n5)
	fmt.Printf("n6的数据类型：%T\n", n6)
	fmt.Printf("n7的数据类型：%T\n", n7)
	fmt.Printf("n8的数据类型：%T\n", n8)
	fmt.Printf("n9的数据类型：%T\n", n9)
	fmt.Printf("n10的数据类型：%T\n", n10)
}
