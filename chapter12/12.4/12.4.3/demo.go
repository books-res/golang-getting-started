package main

import (
	"fmt"
	"reflect"
)

func main() {
	var fds = []reflect.StructField{
		reflect.StructField{
			Name: "Header",
			Type: reflect.TypeOf("zzz"), // string
		},
		reflect.StructField{
			Name: "Data",
			Type: reflect.TypeOf([]byte{}), // []byte
		},
		reflect.StructField{
			Name: "Size",
			Type: reflect.TypeOf(0), // int
		},
		reflect.StructField{
			Name: "Position",
			Type: reflect.PtrTo(reflect.TypeOf(uint(0))), // *uint
		},
	}
	// 创建新的结构体类型
	newStruct := reflect.StructOf(fds)

	fmt.Printf("动态创建的结构体：\n%v\n\n", newStruct)

	// 创建结构体实例
	structVal := reflect.New(newStruct).Elem()
	// 设置字段的值
	structVal.Field(0).SetString("test")
	structVal.Field(1).SetBytes([]byte("c2a5de"))
	structVal.Field(2).SetInt(6)
	var v uint = 60
	structVal.Field(3).Set(reflect.ValueOf(&v))

	fmt.Printf("结构体实例的值：\n%+v\n", structVal)
}
