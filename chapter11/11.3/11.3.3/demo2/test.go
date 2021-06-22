package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var r = ring.New(5)		// 初始化链表实例

	// 给链表中的元素赋值
	// 元素列表：item-1、item-2、item-3、item-4、item-5
	n := r.Len()
	p := r
	for i := 0; i < n; i++ {
		p.Value = fmt.Sprintf("item-%d", i + 1)
		p = p.Next()
	}

	// 滚动链表
	rx := r.Move(-3)
	fmt.Printf("向后滚动3个元素后：%v\n", rx.Value)
}