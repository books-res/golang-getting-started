package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("今天的天气：%s\n", "中雨")

	var a, b, c, d, e = '春', '眠', '不', '觉', '晓'
	fmt.Printf("%c %c %c %c %c\n", a, b, c, d, e)

	var h uint32 = 55660
	// 将32位无符号整数值转换为字符串
	hs := strconv.FormatUint(uint64(h), 10)
	fmt.Printf("幸运数字是：%s\n", hs)
}
