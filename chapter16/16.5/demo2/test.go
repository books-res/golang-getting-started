package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

func main() {
	var (
		// 三个Writer对象
		wt1, _ = os.Create("file-1.txt")
		wt2, _ = os.Create("file-2.txt")
		wt3, _ = os.Create("file-3.txt")
	)
	// 退出时自动关闭文件
	defer wt1.Close()
	defer wt2.Close()
	defer wt3.Close()

	// 生成随机字节
	data := make([]byte, 24)
	rand.Read(data)

	// 写入数据
	writer := io.MultiWriter(wt1, wt2, wt3)
	var n, err = writer.Write(data)
	// 若发生错误，则输出错误信息
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("已成功向三个文件写入%d个字节\n", n)
}
