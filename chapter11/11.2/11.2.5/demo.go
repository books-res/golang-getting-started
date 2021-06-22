package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 初始化双向链表
	var lst = list.New()

	// 向链表添加元素
	lst.PushBack(1000)
	lst.PushBack(2000)
	lst.PushBack(3000)
	lst.PushBack(4000)
	lst.PushBack(5000)
	lst.PushBack(6000)

	// 第一种枚举方法：头 --> 尾
	fmt.Print("从头到尾枚举：")
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v  ", e.Value)
	}
	fmt.Print("\n")

	// 第二种枚举方法：尾 --> 头
	fmt.Print("从尾到头枚举：")
	for e := lst.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v  ", e.Value)
	}
	fmt.Print("\n")
}