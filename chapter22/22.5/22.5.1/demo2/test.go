package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	var cmd = &exec.Cmd{
		Path: "app",
		Dir:  "./ext/bin",
	}
	// 设置环境变量
	cmd.Env = append(cmd.Env, "APP_NAME=Demo")
	cmd.Env = append(cmd.Env, "APP_ID=K786DX-6F-X8")
	cmd.Env = append(cmd.Env, "DEV_ID=7611253")
	// 获取标准输出流的通信管道
	stdOutp, _ := cmd.StdoutPipe()
	// 启动子进程
	err := cmd.Start()
	defer cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取输出内容
	fmt.Print("子进程输出的内容：\n")
	io.Copy(os.Stdout, stdOutp)
}
