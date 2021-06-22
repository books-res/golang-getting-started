package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var r = ring.New(4)

	n := r.Len()	// 链表长度为4
	p := r			// 临时指针
	// 元素列表：1、2、3、4
	for i := 0; i < n; i++ {
		p.Value = i + 1		// 设置元素的值
		p = p.Next()		// 转到下一下元素
	}

	rx := r.Move(18)		// 实际移动 18 % 4 个元素
	fmt.Print(rx.Value)
}