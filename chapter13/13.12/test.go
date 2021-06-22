package main

import (
	"fmt"
	"strings"
)

func main() {
	// 大写 -> 小写
	var x = "CXPKH"
	var y = strings.ToLower(x)
	fmt.Printf("%q转换为小写后：%v\n", x, y)

	// 小写 -> 大写
	x = "capture"
	y = strings.ToUpper(x)
	fmt.Printf("%q转换为大写后：%v\n", x, y)

	// 不起作用
	x = "前进123"
	y = strings.ToUpper(x)
	fmt.Printf("%q转换为大写后：%v\n", x, y)

	// 同样有效
	x = "ΩЙТЦ"
	y = strings.ToLower(x)
	fmt.Printf("%q转换为小写后：%v\n", x, y)
}
