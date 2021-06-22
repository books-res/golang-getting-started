package main

import (
	"fmt"
	"reflect"
)

type dog struct {
	Nick string
	Color string
	Age uint8
}

func main() {
	var mypet = dog{
		Nick: "Peter",
		Color: "black",
		Age: 2,
	}

	fmt.Println("----- 修改前 -----")
	printValues(mypet)

	setValue(&mypet, "Age", uint8(5))

	fmt.Println("\n----- 修改后 -----")
	printValues(mypet)
}

func printValues(obj interface{}) {
	var theVal = reflect.ValueOf(obj)
	// 如果传递过来的是对象的指针
	// 那么先获取该指针所指向的对象
	if theVal.Kind() == reflect.Ptr {
		theVal = theVal.Elem()
	}
	// 获取字段成员数量
	ln := theVal.NumField()
	// 访问所有字段
	for i := 0; i < ln; i++ {
		tm := theVal.Type().Field(i).Name
		vm := theVal.Field(i).Interface()
		fmt.Printf("%s: %v\n", tm, vm)
	}
}

func setValue(obj interface{}, fdname string, fdval interface{}) {
	var objval = reflect.ValueOf(obj)
	if objval.Kind() != reflect.Ptr {
		fmt.Println("请使用指针类型")
		return
	}
	objval = objval.Elem()
	// 查找字段
	fd := objval.FieldByName(fdname)
	if fd.IsValid() == false {
		fmt.Println("未找到目标字段")
		return
	}
	// 验证一下值的类型是否与字段匹配
	newVal := reflect.ValueOf(fdval)
	if fd.Kind() != newVal.Kind() {
		fmt.Println("字段值的类型不匹配")
		return
	}
	// 设置新值
	fd.Set(newVal)
}