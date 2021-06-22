package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		C1 chan uint
		C2 chan <- bool
	)

	t1 := reflect.TypeOf(C1)
	fmt.Println("----- C1 变量 -----")
	fmt.Printf("通道类型：%s\n", t1.Name())
	fmt.Print("通信方向：")
	switch t1.ChanDir() {
	case reflect.RecvDir:
		fmt.Print("只能从通道接收数据")
	case reflect.SendDir:
		fmt.Print("只能向通道发送数据")
	case reflect.BothDir:
		fmt.Print("同时支持发送和接收数据")
	}

	t2 := reflect.TypeOf(C2)
	fmt.Println("\n\n----- C2 变量 -----")
	fmt.Printf("通道类型：%s\n", t2.Name())
	fmt.Print("通信方向：")
	switch t2.ChanDir() {
	case reflect.RecvDir:
		fmt.Print("只能从通道接收数据")
	case reflect.SendDir:
		fmt.Print("只能向通道发送数据")
	case reflect.BothDir:
		fmt.Print("同时支持发送和接收数据")
	}
}
