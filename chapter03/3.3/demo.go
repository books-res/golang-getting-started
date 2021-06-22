package main

import (
	"fmt"
)

func main() {
	var result bool
	// 1、比较两个整值
	var a, b = 85, 87
	result = a > b
	fmt.Printf("%d大于%d吗？%t\n", a, b, result)

	// 2、比较两个浮点数
	var e, f = 0.77, 2.565
	result = e < f
	fmt.Printf("%.3f小于%.3f吗？%t\n", e, f, result)

	// 3、比较两个字符串
	var str1 = "yes"
	var str2 = "Yes"
	result = str1 == str2
	fmt.Printf("“%s”与“%s”是否相等？%t\n", str1, str2, result)

	// 4、比较指针类型
	var num = 30000
	var p1 *int = &num
	var p2 *int = &num
	result = p1 == p2
	fmt.Printf("p1与p2是否相等？%t\n", result)
	// 比较不同类型的指针
	// 以下代码会出错
	// var p3 *string = new(string)
	// var p4 *float32 = new(float32)
	// result = p3 == p4

	// 5、声明为interface的变量比较
	var k1, k2 interface{}
	k1 = 50
	k2 = 50
	result = k1 == k2
	fmt.Printf("k1与k2是否相等？%t\n", result)

	// 6、比较两个结构体实例
	// 声明变量并赋值
	var obj1, obj2 test
	obj1 = test{
		fd01: 2,
		fd02: 0.00075,
		fd03: 809,
	}
	obj2 = test{
		fd01: 2,
		fd02: 0.00075,
		fd03: 809,
	}
	result = obj1 == obj2
	fmt.Printf("obj1与obj2是否相等？%t\n", result)

	// 7、接口变量比较
	var (
		pet1 pet = cat{name: "Jack"}
		pet2 pet = dog{name: "Tim"}
		pet3 pet = pet2
	)
	result = pet1 == pet2
	fmt.Printf("pet1与pet2是否相等？%t\n", result)
	result = pet2 == pet3
	fmt.Printf("pet2与pet3是否相等？%t", result)
}

/*---------------------------------------*/
// 定义新的结构体
type test struct {
	fd01 int8
	fd02 float64
	fd03 int16
}

/*---------------------------------------*/
// 定义新接口
type pet interface {
	GetName() string
}

type cat struct {
	name string
}

func (c cat) GetName() string {
	return c.name
}

type dog struct {
	name string
}

func (d dog) GetName() string {
	return d.name
}
