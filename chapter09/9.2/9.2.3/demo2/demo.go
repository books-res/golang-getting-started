package main

import "fmt"

type demo struct {
	data int
}

func (x demo) setIntV1(n int) {
	x.data = n		// 修改data字段的值
}

func (x *demo) setIntV2(n int) {
	x.data = n		// 修改data字段的值
}

func main() {
	// 初始化demo实例
	var a = demo{ data: 100 }

	// 情况一：非指针类型接收demo实例
	fmt.Println("--------- 传递demo实例的副本 ---------")
	fmt.Printf("调用setIntV1方法前，data字段的值：%d\n", a.data)
	// 调用setIntV1方法
	a.setIntV1(200)
	fmt.Printf("调用setIntV1方法后，data字段的值：%d\n\n", a.data)

	// 情况二：以指针类型接收demo实例
	fmt.Println("--------- 传递demo实例的内存地址 ---------")
	fmt.Printf("调用setIntV2方法前，data字段的值：%d\n", a.data)
	// 调用setIntV2方法
	a.setIntV2(200)
	fmt.Printf("调用setIntV2方法后，data字段的值：%d\n", a.data)
}