package main

import (
	"fmt"
	"os"
)

func main() {
	const fn = "demo.txt" // 文件名
	// 打开文件
	// O_APPEND：追加内容
	// O_CREATE：如果文件不存在，则创建它
	// 权限：0244
	// 所有者：写入权限；用户组：读取权限；其他用户：读取权限
	file, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE, 0244)
	if err != nil {
		fmt.Println(err)
		return
	}

	var input string
	fmt.Print("请输入要写入到文件的内容：\n")
	// 读取键盘输入的内容
	fmt.Scan(&input)
	// 将内容写入文件
	file.WriteString(input + "\n")
	// 关闭文件
	file.Close()
	fmt.Println("文件写入成功")
}
