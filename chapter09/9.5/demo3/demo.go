package main

import "fmt"

// 接口
type test1 interface {
	readFromFile() []byte
}
type test2 interface {
	sendEmail(to, subject string, body []byte) int
}
type test3 interface {
	getResponse() (int, []byte)
}

// 分别实现以上接口
type demoType1 struct { }
func (t demoType1) readFromFile() []byte {
	return []byte("kjidfdf56gg566")
}

type demoType2 struct { }
func (t demoType2) sendEmail(to, subject string, body []byte) int {
	return 1
}

type demoType3 struct { }
func (t demoType3) getResponse() (int, []byte) {
	return 200, []byte("e2068xz4yb7owelk9sye")
}

func main() {
	var (
		a interface{} = demoType1{}
		b interface{} = demoType2{}
		c interface{} = demoType3{}
	)

	
	if x, ok := a.(test1); ok {
		fmt.Printf("变量a的实际类型为%T，实现了test1接口\n", x)
	}
	
	if x, ok := b.(test2); ok {
		fmt.Printf("变量b的实际类型为%T，实现了test2接口\n", x)
	}

	if x, ok := c.(test3); ok {
		fmt.Printf("变量c的实际类型为%T，实现了test3接口\n", x)
	}
}