package main

import "fmt"

// 空接口
var a interface{}

// 空结构体
var b struct{}

func main() {
	var obj interface{}
	obj = 12345
	fmt.Printf("%+-20[1]v%[1]T\n", obj)
	obj = "xyz"
	fmt.Printf("%+-20[1]v%[1]T\n", obj)
	obj = 3.1415927
	fmt.Printf("%+-20[1]v%[1]T\n", obj)
	obj = struct{E int; S string} {200, "opqrst"}
	fmt.Printf("%+[1]v%30[1]T\n", obj)
}