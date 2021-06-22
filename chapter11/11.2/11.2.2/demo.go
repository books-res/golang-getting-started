package main

import (
	"container/list"
	"fmt"
)

func main() {
	var mylist = list.New()
	fmt.Printf("新链表实例的内存地址：%p\n", mylist)
}