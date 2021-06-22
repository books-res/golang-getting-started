package main

import (
	"fmt"
	"net"
)

func main() {
	// 获取某个网络接口
	var wlanInf, err = net.InterfaceByName("WLAN")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 列出与此接口关联的地址
	fmt.Printf("与网络接口%s关联的单播地址列表：\n", wlanInf.Name)
	uaddrs, _ := wlanInf.Addrs()
	for _, a := range uaddrs {
		fmt.Printf("\t%s\n", a)
	}
	fmt.Printf("\n与网络接口%s关联的多播地址列表：\n", wlanInf.Name)
	mcaddrs, _ := wlanInf.MulticastAddrs()
	for _, a := range mcaddrs {
		fmt.Printf("\t%s\n", a)
	}
}
