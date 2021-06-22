package main

import (
	"fmt"
	"net"
)

func main() {
	// 通过编号获取网络接口信息
	var wanInt, err = net.InterfaceByIndex(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印接口信息
	fmt.Printf("接口名称：%s\n", wanInt.Name)
	fmt.Printf("最大传输单元（MTU）：%d\n", wanInt.MTU)
}
