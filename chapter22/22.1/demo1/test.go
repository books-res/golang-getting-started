package main

import (
	"fmt"
	"net"
)

func main() {
	interFaceList, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印各个网络接口的信息
	for _, f := range interFaceList {
		fmt.Printf("接口编号：%d\n", f.Index)
		fmt.Printf("接口名称：%s\n", f.Name)
		fmt.Printf("最大传输单元（MTU）：%d\n", f.MTU)
		fmt.Printf("硬件地址：%x\n", f.HardwareAddr)
		fmt.Print("\n")
	}
}
