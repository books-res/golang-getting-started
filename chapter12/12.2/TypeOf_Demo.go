package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var a uint16 = 2800
	var tp reflect.Type = reflect.TypeOf(a)
	fmt.Printf("类型名称：%s\n", tp.Name())
	fmt.Printf("占用内存大小（字节）：%d\n", tp.Size())
	fmt.Printf("占用内存大小（位）：%d\n", tp.Bits())
}