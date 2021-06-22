package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 四个Reader对象
	var (
		rd1 = strings.NewReader("第一个数据源\n")
		rd2 = strings.NewReader("第二个数据源\n")
		rd3 = strings.NewReader("第三个数据源\n")
		rd4 = strings.NewReader("第四个数据源\n")
	)

	var mtRd = io.MultiReader(rd1, rd2, rd3, rd4)

	fmt.Println("------ 从多个源中读到的内容 ------")
	// 将数据复制到标准输出流
	io.Copy(os.Stdout, mtRd)
}
