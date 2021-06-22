package main

import (
	"unsafe"
	"fmt"
)

func main() {
	// 无符号整数
	var (
		n1 uint8 = 122
		n2 uint16 = 2000
		n3 uint32 = 53530020
		n4 uint64 = 4.33e+5
		n5 uint = 99977723
	)
	fmt.Printf("8位无符号整数：%d\n", unsafe.Sizeof(n1))
	fmt.Printf("16位无符号整数：%d\n", unsafe.Sizeof(n2))
	fmt.Printf("32位无符号整数：%d\n", unsafe.Sizeof(n3))
	fmt.Printf("64位无符号整数：%d\n", unsafe.Sizeof(n4))
	fmt.Printf("无符号整数（uint）：%d\n", unsafe.Sizeof(n5))

	// 有符号整数
	var (
		s1 int8 = 103
		s2 int16 = 999
		s3 int32 = 120020
		s4 int64 = 65630000
		s5 int = 7.323e+7
	)
	fmt.Printf("8位有符号整数：%d\n", unsafe.Sizeof(s1))
	fmt.Printf("16位有符号整数：%d\n", unsafe.Sizeof(s2))
	fmt.Printf("32位有符号整数：%d\n", unsafe.Sizeof(s3))
	fmt.Printf("64位有符号整数：%d\n", unsafe.Sizeof(s4))
	fmt.Printf("有符号整数（int）：%d\n", unsafe.Sizeof(s5))

	// 浮点数
	var (
		f1 float32 = 0.0005909
		f2 float64 = 1.10000000012525
	)
	fmt.Printf("32位浮点数：%d\n", unsafe.Sizeof(f1))
	fmt.Printf("64位浮点数：%d\n", unsafe.Sizeof(f2))

	// 复数
	var (
		c1 complex64 = 19 + 7i
		c2 complex128 = -3 + 0.6i
	)
	fmt.Printf("64位复数：%d\n", unsafe.Sizeof(c1))
	fmt.Printf("128位复数：%d\n", unsafe.Sizeof(c2))
}