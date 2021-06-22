package main

import "fmt"

// 自定义接口类型
type tester interface {
	getDescription() string
}

// 自定义结构体
type data1 struct { }
func (d data1) getDescription() string {
	return "Data v1"
}

type data2 struct { }
func (d data2) getDescription() string {
	return "Data v2"
}

type data3 struct { }
func (d data3) getDescription() string {
	return "Data v3"
}

func main() {
	// 必须声明为接口类型
	var s tester = data3{ }

	switch val := s.(type) {
	case data1:
		fmt.Println(val.getDescription())
	case data2:
		fmt.Println(val.getDescription())
	case data3:
		fmt.Println(val.getDescription())
	}
}