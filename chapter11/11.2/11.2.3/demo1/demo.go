package main

import (
	"container/list"
	"fmt"
)

func printElems(ls *list.List) {
	for x := ls.Front(); x != nil; x = x.Next() {
		fmt.Printf("%v  ", x.Value)
	}
}

func main() {
	// 初始化链表对象
	var mylist = list.New()
	// 添加一个元素
	mylist.PushFront(100)
	// 在末尾追加三个元素
	mylist.PushBack(200)
	mylist.PushBack(300)
	mylist.PushBack(400)
	fmt.Print("此时链表中有四个元素：")
	printElems(mylist)

	// 在链表首部插入一个元素
	mylist.PushFront(90)
	fmt.Print("\n在首部插入元素：")
	printElems(mylist)

	// 在倒数第二个元素前插入元素
	last := mylist.Back()	// 最后一个元素
	x := last.Prev()		// 倒数第二个元素
	mylist.InsertBefore(700, x)
	fmt.Print("\n在倒数第二个元素前插入元素：")
	printElems(mylist)

	// 在第一个元素之后插入元素
	first := mylist.Front()		// 第一个元素
	mylist.InsertAfter(800, first)
	fmt.Print("\n在第一个元素后插入元素：")
	printElems(mylist)
}