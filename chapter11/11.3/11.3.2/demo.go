package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建实例
	var myring = ring.New(5)
	// 设置元素值
	n := myring.Len()		// 元素个数
	pt := myring			// 临时指针
	v := 'A'
	for x := 0; x < n; x++ {
		pt.Value = v
		pt = pt.Next()
		v++
	}
	// 循环输出
	pt = myring
	for n := 0; n < 15; n++ {
		fmt.Printf("%c ", pt.Value)
		pt = pt.Next()
	}
}