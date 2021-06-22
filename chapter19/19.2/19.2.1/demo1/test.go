package main

import (
	"flag"
	"fmt"
)

func main() {
	// 注册命令行参数
	pWid := flag.Int("width", 0, "矩形的宽度")
	pHei := flag.Int("height", 0, "矩形的高度")
	// 重要：一定要调用此函数
	flag.Parse()

	// 计算面积
	ar := (*pWid) * (*pHei)
	fmt.Printf("矩形的面积为：%d\n", ar)
}
