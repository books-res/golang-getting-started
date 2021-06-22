package main

import "fmt"

func main() {
	type cat struct {
		name string
		age  int
	}

	// 实例化cat
	var c = cat{"Jim", 3}
	fmt.Printf("用%%v控制符输出cat对象：%v\n", c)
	fmt.Printf("用%%+v控制符输出cat对象：%+v\n", c)
	fmt.Printf("用%%#v控制符输出cat对象：%#v\n", c)
}
