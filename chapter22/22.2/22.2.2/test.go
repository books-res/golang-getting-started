package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	// 命令行参数
	var (
		// 程序角色，服务器或客户端
		role string
		// 相关的地址
		// 若角色为服务器，则addr为侦听地址
		// 若角色为客户端，则addr为连接的服务器地址
		addr string
	)
	// 将变量与命令行参数绑定
	flag.StringVar(&role, "role", "server", "程序角色。可选值为 server 或 client")
	flag.StringVar(&addr, "addr", ":8733", "若角色为server，则此地址为侦听地址；若角色为client，则此地址为要连接的服务器地址")

	// 分析命令参数
	flag.Parse()

	if strings.ToLower(role) == "server" {
		startServer(addr)
	} else if strings.ToLower(role) == "client" {
		startClient(addr)
	} else {
		fmt.Println("程序角色无效")
	}
}

func startServer(listenAddr string) {
	fmt.Println("程序的当前角色：Server")
	var udpaddr, err = net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	udpc, err := net.ListenUDP("udp", udpaddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer udpc.Close()
	fmt.Printf("服务器侦听地址：%s\n", udpc.LocalAddr())
	fmt.Print("开始接收消息……\n\n")
	// 循环接收消息
	for {
		reader := bufio.NewReader(udpc)
		msg, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
			continue
		}
		// 将接收到的消息打印到屏幕上
		fmt.Printf(">>> %s", msg)
	}
}

func startClient(svaddr string) {
	fmt.Println("程序的当前角色：Client")
	udpsvaddr, err := net.ResolveUDPAddr("udp", svaddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 发起连接
	udpc, err := net.DialUDP("udp", nil, udpsvaddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("准备就绪，你可以输入要发送的消息")
	fmt.Print("【输入end可退出】\n\n")

	var inputMsg string
	for {
		fmt.Print("请输入：")
		stdReader := bufio.NewReader(os.Stdin)
		inputMsg, _ = stdReader.ReadString('\n')
		chkend := strings.ToLower(inputMsg)
		// 去掉不需要的字符
		chkend = strings.Trim(chkend, "\r\n\t ")
		// 如果输入的消息为end，则跳出循环
		if chkend == "end" {
			break
		}
		udpc.Write([]byte(inputMsg))
	}
	// 关闭连接
	udpc.Close()
}
