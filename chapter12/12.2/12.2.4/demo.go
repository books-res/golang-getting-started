package main

import (
	"fmt"
	"reflect"
)

type address struct {
	Province string
	City string
	Town string
}

type person struct {
	Name string
	Age uint8
	address
}

type employee struct {
	person
	Department string
	Code uint64
}

func main() {
	var emp employee
	// 获取类型信息
	tp := reflect.TypeOf(emp)
	// 查找Age字段
	fdName,ok := tp.FieldByName("Age")
	if ok {
		fmt.Printf("%s字段的索引：%v\n", fdName.Name, fdName.Index)
	}
	// 查找City字段
	fdCity,ok := tp.FieldByName("City")
	if ok {
		fmt.Printf("%s字段的索引：%v\n", fdCity.Name, fdCity.Index)
	}
	// 直接通过索引查找
	theIndex := []int{ 0, 2, 2 }
	fdTown := tp.FieldByIndex(theIndex)
	fmt.Printf("\n索引为%v的字段信息：\n", theIndex)
	fmt.Printf("字段名称：%s\n", fdTown.Name)
	fmt.Printf("字段类型：%v\n", fdTown.Type)
}