package main

import (
	"container/list"
	"fmt"
)

func main() {

	// 初始化链表
	var ls = list.New()
	
	// 添加五个元素
	ls.PushBack(1)
	ls.PushBack(2)
	ls.PushBack(3)
	ls.PushBack(4)
	ls.PushBack(5)

	fmt.Print("原顺序：")
	printList(ls)

	// 获取第一个、最后一个元素的引用
	first, last := ls.Front(), ls.Back()

	// 获取第二个、第四个元素的引用
	second, fourth := first.Next(), last.Prev()

	// 第二个元素移到第四位
	ls.MoveBefore(second, last)
	// 第四个元素移到第二位
	ls.MoveAfter(fourth, first)

	// 将第一个元素和最后一个元素的位置互换
	ls.MoveToFront(last)
	ls.MoveToBack(first)

	fmt.Print("移动后：")
	printList(ls)
}

func printList(x *list.List) {
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v  ", e.Value)
	}
	fmt.Println()
}