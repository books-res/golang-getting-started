package main

import (
	"fmt"
	"reflect"
)

type photoPlayer struct { }
func (x photoPlayer) Start() { }
func (x photoPlayer) Stop(isclosing bool) int { return 0 }

func main() {
	var obj = photoPlayer{ }
	// 获取结构体的方法列表
	ty := reflect.TypeOf(obj)
	fmt.Printf("类型：%s\n\n", ty.Name())
	// 获取接口中方法的数量
	var mn = ty.NumMethod()
	for i := 0; i < mn; i++ {
		mt := ty.Method(i)
		fmt.Printf("方法名称：%s\n", mt.Name)
		fmt.Printf("方法的函数类型：%v\n", mt.Type)
	}
}