package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建新文件
	file, err := os.Create("abc.txt")
	// 判断是否发生错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 向文件写入三行文本
	fmt.Fprintln(file, "第一行")
	fmt.Fprintln(file, "第二行")
	fmt.Fprintln(file, "第三行")

	// 向文件写入三个不同类型的值
	var x, y, z = -2000, 0.25, 12 + 4i
	fmt.Fprint(file, x, y, z)
}
