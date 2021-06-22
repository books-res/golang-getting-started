package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 初始化双向链表实例
	theList := list.New()
	// 添加四个元素
	e1 := theList.PushBack("a-b-c-d")
	e2 := theList.PushBack("h-i-j-k")
	e3 := theList.PushBack(-5000)
	e4 := theList.PushBack(-3000)
	// 即将删除元素
	fmt.Printf("链表中已存元素个数：%d\n", theList.Len())

	// 删除第一个元素
	val1 := theList.Remove(e1)
	fmt.Printf("删除【%v】后，还剩%d个元素\n", val1, theList.Len())

	// 删除第二个元素
	val2 := theList.Remove(e2)
	fmt.Printf("删除【%v】后，还剩%d个元素\n", val2, theList.Len())

	// 删除第三个元素
	val3 := theList.Remove(e3)
	fmt.Printf("删除【%v】后，还剩%d个元素\n", val3, theList.Len())

	// 删除第四个元素
	val4 := theList.Remove(e4)
	fmt.Printf("删除【%v】后，还剩%d个元素\n", val4, theList.Len())
}