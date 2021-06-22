package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	var newcmd = &exec.Cmd{
		Path: "app",
		Dir:  "./ext/bin",
	}
	// 在子进程启动前获取标准输入/输出流的访问管道
	stdIn, _ := newcmd.StdinPipe()
	stdOut, _ := newcmd.StdoutPipe()
	// 启动子进程
	err := newcmd.Start()
	// 退出该范围前等待子进程结束
	defer newcmd.Wait()

	if err != nil {
		fmt.Println(err)
		return
	}
	// 用标准输入流写入内容
	stdIn.Write([]byte("1.5 3.6"))
	stdIn.Close()
	// 读取结果
	io.Copy(os.Stdout, stdOut)
	stdOut.Close()
}
