package main

import "fmt"

type test struct {
	xn int16
	xt float32
}

type test2 struct {
	t2 test
}

func main() {
	// 声明变量
	var v test
	//fmt.Printf("%+v\n", v)
	// 给字段赋值
	v.xn = 900
	v.xt = 17.00061

	fmt.Printf("%+v\n", v)

	//------------------------------
	var v2 test2
	v2.t2.xn = 450
	v2.t2.xt = 0.33
	fmt.Printf("%+v", v2)
}