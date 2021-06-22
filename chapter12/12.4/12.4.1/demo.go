package main

import (
	"fmt"
	"reflect"
)

func main() {
	var v1 = reflect.New(reflect.TypeOf(1))
	fmt.Printf("v1的类型：%v\n", v1.Type())
	var v2 = reflect.New(reflect.TypeOf(struct {
		F1 uint
		F2 string
	}{}))
	fmt.Printf("v2的类型：%v\n", v2.Type())
	var v3 = reflect.New(reflect.TypeOf(func(x string) {}))
	fmt.Printf("v3的类型：%v\n", v3.Type())
}
