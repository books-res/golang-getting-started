package main

import "fmt"

func main() {
	var px *int16 = new(int16)
	fmt.Printf("px指向的内存地址：%p\n", px)
	fmt.Printf("px所指向对象的值：%v", *px)

	var py = new(string)
	fmt.Printf("py是否为nil？ %t\n", py == nil)

	// 为结构体分配内存空间后
	// a、b字段将以其所属类型的默认值来初始化
	var ptr = new(demo)
	fmt.Printf("ptr.a: %v\n", ptr.a)
	fmt.Printf("ptr.b: %v\n", ptr.b)

	// 接口类型默认分配nil
	var pi = new(sender)
	fmt.Printf("pi指向的内存地址：%p\n", pi)
	fmt.Printf("pi所指向内存对应的值：%v\n", *pi)
}

type demo struct {
	a int
	b bool
}

type sender interface {
	Lineout(data []byte) int
}