package main

import "fmt"

type base struct {
	code uint
	line uint64
	label string
}

type dev struct {
	base				// 嵌套
	size float32
	publisher string
}

func main() {
	var x dev //= dev{base{1001,1,"F7"},1.337,"Jack"}
	x.code = 1001			// 等同于x.base.code
	x.line = 1				// 等同于x.base.line
	x.label = "F7"			// 等同于x.base.label
	x.size = 1.337
	x.publisher = "Jack"
	fmt.Printf("code: %d\n", x.code)
	fmt.Printf("label: %s\n", x.label)
	fmt.Printf("line: %d\n", x.line)
	fmt.Printf("size: %f\n", x.size)
	fmt.Printf("publisher: %s\n", x.publisher)

	fmt.Println()
	//var y = dev{ base: base{1002,1,"D6"}, size: 0.12, publisher: "Dick"}
	var y = dev{ base{1002,1,"D6"}, 0.12, "Dick"}
	fmt.Printf("code: %d\n", y.code)
	fmt.Printf("line: %d\n", y.line)
	fmt.Printf("label: %s\n", y.label)
	fmt.Printf("size: %f\n", y.size)
	fmt.Printf("publisher: %s\n", y.publisher)

	// 下面代码会出错
	//var z base = dev{base{1003,1,"C8"}, 0.025, "July"}
	//fmt.Printf("%+v\n", z)
}