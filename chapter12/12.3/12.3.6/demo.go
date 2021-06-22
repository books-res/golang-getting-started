package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 创建映射实例
	var mp = make(map[string]float32)

	// 获取相关的Value对象
	valOfMap := reflect.ValueOf(mp)
	// 设置映射元素
	valOfMap.SetMapIndex(
		reflect.ValueOf("T1"),            // key
		reflect.ValueOf(float32(0.0071)), // value
	)
	valOfMap.SetMapIndex(
		reflect.ValueOf("T2"),           // key
		reflect.ValueOf(float32(9.202)), // value
	)
	valOfMap.SetMapIndex(
		reflect.ValueOf("T3"),           // key
		reflect.ValueOf(float32(-0.03)), // value
	)

	// 获取key集合
	var keys = valOfMap.MapKeys()
	// 输出映射中的元素列表
	for _, k := range keys {
		v := valOfMap.MapIndex(k)
		fmt.Printf("%v - %v\n", k.Interface(), v.Interface())
	}
}
