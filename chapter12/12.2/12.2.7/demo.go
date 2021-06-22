package main

import (
	"fmt"
	"reflect"
)

// 定义接口
type ball interface {
	Play()
}
type gymnastics interface {
	Do()
}

type footBall struct { }
func (b footBall) Play() {
	fmt.Println("踢足球")
}

func main() {
	var (
		// 获取ball接口类型的Type
		tpofball = reflect.TypeOf(new(ball)).Elem()
		// 获取gymnastics接口类型的Type
		tpofgy = reflect.TypeOf(new(gymnastics)).Elem()
		// 获取footBall结构体的Type
		tpoffb = reflect.TypeOf(footBall{})
	)

	// 情形一：是否实现了ball接口
	b := tpoffb.Implements(tpofball)
	fmt.Printf("结构体%s", tpoffb.Name())
	if b {
		fmt.Print("实现了")
	} else {
		fmt.Print("未实现")
	}
	fmt.Printf("%s接口\n", tpofball.Name())

	// 情形二：是否实现了gymnastics接口
	b = tpoffb.Implements(tpofgy)
	fmt.Printf("结构体%s", tpoffb.Name())
	if b {
		fmt.Print("实现了")
	} else {
		fmt.Print("未实现")
	}
	fmt.Printf("%s接口\n", tpofgy.Name())
}