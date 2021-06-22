package main

import (
	"container/list"
	"fmt"
)

func main() {
	var myls = list.New()
	// 添加五个元素
	// -1、-2、-3、-4、-5
	myls.PushFront(-5)
	myls.PushFront(-4)
	myls.PushFront(-3)
	myls.PushFront(-2)
	myls.PushFront(-1)

	// 从前往后删除
	/*
	for e := myls.Front(); e != nil; e = myls.Front() {
		tmp := myls.Remove(e)
		fmt.Printf("已删除 %v\n", tmp)
	}
	*/

	// 从后向前删除
	for e := myls.Back(); e != nil; e = myls.Back() {
		tmp := myls.Remove(e)
		fmt.Printf("已删除 %v\n", tmp)
	}
}