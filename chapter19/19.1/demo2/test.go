package main

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func main() {
	// 本例仅需要一个命令行参数
	// 加上程序名称，Args中应有两个元素
	if len(os.Args) != 2 {
		fmt.Println("传递的参数过多")
		return
	}
	// 获取要做SHA1运算的内容
	var content = os.Args[1]
	var res = sha1.Sum([]byte(content))
	// 输出结果
	fmt.Printf("sha1：%x\n", res)
}
