package main

import (
	"fmt"
	"reflect"
)

func main() {
	var kx = []int{1, 5, 7, 12}
	fmt.Println("更新前：", kx)
	updateElement(kx, 1, 9000)
	fmt.Println("更新后：", kx)
}

func updateElement(obj interface{}, index int, elval interface{}) {
	v := reflect.ValueOf(obj)
	// 要求目标对象是数组或者切片类型
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		// 获取元素的总数量
		ln := v.Len()
		// 验证指定的索引是否有效
		if index < 0 || index >= ln {
			fmt.Println("索引超出有效范围")
			return
		}
		// 获取元素值
		oldval := v.Index(index)
		newval := reflect.ValueOf(elval)
		// 验证新值的类型是否与旧值匹配
		if oldval.Kind() != newval.Kind() {
			fmt.Println("元素值的类型不匹配")
			return
		}
		// 更新元素值
		oldval.Set(newval)
	} else {
		fmt.Println("对象类型不是数组或切片类型")
		return
	}
}
