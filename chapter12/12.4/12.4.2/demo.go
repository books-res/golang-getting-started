package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 元素类型
	var elemtp = reflect.TypeOf("xxx")
	// 构建数组类型
	var arrtype = reflect.ArrayOf(5, elemtp)
	// 实例化数组
	var arrVal = reflect.New(arrtype).Elem()
	// 设置元素值
	arrVal.Index(0).SetString("Item_1")
	arrVal.Index(1).SetString("Item_2")
	arrVal.Index(2).SetString("Item_3")
	arrVal.Index(3).SetString("Item_4")
	arrVal.Index(4).SetString("Item_5")
	// 获取真实的数组对象
	theArr, ok := arrVal.Interface().([5]string)
	if ok {
		// 枚举元素
		for i, e := range theArr {
			fmt.Printf("%d - %s\n", i, e)
		}
	}
}
