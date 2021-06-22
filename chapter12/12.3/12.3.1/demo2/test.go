package main

import (
	"fmt"
	"reflect"
)

func main() {
	var bv = false
	fmt.Printf("变量的原值：%v\n", bv)

	var val = reflect.ValueOf(&bv)
	// 获取指针指向的值
	var blval = val.Elem()
	if blval.Kind() == reflect.Bool && blval.CanSet() {
		blval.SetBool(true)
	}

	fmt.Printf("变量的最新值：%v\n", bv)
}