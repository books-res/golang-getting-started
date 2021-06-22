package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	id uint
	name, city string
	course string
}

// 读写id字段
func (x Student) GetID() uint {
	return x.id
}
func (x *Student) SetID(id uint) {
	x.id = id
}

// 读写name字段
func (x Student) GetName() string {
	return x.name
}
func (x *Student) SetName(name string) {
	x.name = name
}

// 读写city字段
func (x Student) GetCity() string {
	return x.city
}
func (x *Student) SetCity(city string) {
	x.city = city
}

// 读写course字段
func (x Student) GetCourse() string {
	return x.course
}
func (x *Student) SetCourse(course string) {
	x.course = course
}

//--------------------------------------
func main() {
	var obj Student

	// 用反射技术设置字段的值
	valOfObj := reflect.ValueOf(&obj)
	//fmt.Printf("方法数量：%d\n", valOfObj.NumMethod())
	// 调用SetID方法
	m := valOfObj.MethodByName("SetID")
	p := reflect.ValueOf(uint(187005))
	m.Call([]reflect.Value{ p })
	// 调用SetName方法
	m = valOfObj.MethodByName("SetName")
	p = reflect.ValueOf("小刘")
	m.Call([]reflect.Value{ p })
	// 调用SetCity方法
	m = valOfObj.MethodByName("SetCity")
	p = reflect.ValueOf("珠海")
	m.Call([]reflect.Value{ p })
	// 调用SetCourse方法
	m = valOfObj.MethodByName("SetCourse")
	p = reflect.ValueOf("C++")
	m.Call([]reflect.Value{ p })

	// 输出各字段的值
	fmt.Println("使用反射设置字段的值后：")
	fmt.Printf("id: %v\n", obj.GetID())
	fmt.Printf("name: %v\n", obj.GetName())
	fmt.Printf("city: %v\n", obj.GetCity())
	fmt.Printf("course: %v\n", obj.GetCourse())
}