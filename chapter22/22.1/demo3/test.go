package main

import (
	"fmt"
	"net"
)

func main() {
	// 根据名称获取网络接口
	btInterface, err := net.InterfaceByName("蓝牙网络连接")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("接口名称：%s\n", btInterface.Name)
}
