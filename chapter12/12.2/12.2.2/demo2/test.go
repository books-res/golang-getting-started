package main

import (
	"fmt"
	"reflect"
)

type keyInput struct {
	keycode int32
}
func (ki keyInput) GetKeycode() int32 {
	return ki.keycode
}
func (ki keyInput) SendKey(key int32) { }

func main() {
	var obj keyInput
	// 获取Type对象
	thetp := reflect.TypeOf(obj)

	// 查找名为SendKey的方法
	var methodTofind = "SendKey"
	m, ok := thetp.MethodByName(methodTofind)
	if ok {
		// 如果找到就输出方法的信息
		fmt.Printf("方法%s的函数签名：%v\n", m.Name, m.Type)
	} else {
		// 如果找不到，输出错误信息
		fmt.Printf("%s类型中不存在%s方法\n", thetp.Name(), methodTofind)
	}

	// 查找名为Copy的方法
	methodTofind = "Copy"
	m, ok = thetp.MethodByName(methodTofind)
	if ok {
		fmt.Printf("方法%s的函数签名：%v\n", m.Name, m.Type)
	} else {
		fmt.Printf("%s类型中不存在%s方法", thetp.Name(), methodTofind)
	}
}