package main

import (
	"fmt"
	"reflect"
)

func main() {
	var who string = "小明"
	fmt.Printf("变量who的原值：%v\n", who)

	// 通过反射技术修改变量的值
	val := reflect.ValueOf(&who)
	// 获取指针指向的对象
	val = val.Elem()
	if val.Kind() == reflect.String {
		if val.CanSet() {
			val.Set(reflect.ValueOf("小吴"))
		} else {
			fmt.Println("变量who不允许修改")
		}
	}

	fmt.Printf("修改后，变量who的值：%v\n", who)
}