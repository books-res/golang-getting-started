package main

import "fmt"

// 定义新的结构体
type test struct {
	tag string
}
// test的方法
func (o *test) setTag(v string) {
	o.tag = v
}

func main() {
	var s test = test{ tag: "abc" }
	fmt.Printf("调用setTag方法前，tag字段的值：%s\n", s.tag)
	s.setTag("xyz")
	fmt.Printf("调用setTag方法后，tag字段的值：%s\n", s.tag)
}