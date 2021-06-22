package main

import (
	"fmt"
	"strings"
)

func main() {
	var text = "明日复明日，明日何其多。我生待明日，万事成蹉跎"
	// 查找“明日”第一次出现的位置
	index1 := strings.Index(text, "明日")
	index1 = runeIndexOf(text, index1)
	fmt.Printf("“明日”第一次出现的位置：%d\n", index1)
	// 查找“明日”最后一次出现的位置
	index2 := strings.LastIndex(text, "明日")
	index2 = runeIndexOf(text, index2)
	fmt.Printf("“明日”最后一次出现的位置：%d\n", index2)
	// 查找“昨日”在字符串中的位置
	index3 := strings.Index(text, "昨日")
	index3 = runeIndexOf(text, index3)
	fmt.Printf("“昨日”在字符串中的位置：%d\n", index3)
}

func runeIndexOf(src string, byteIndex int) int {
	if byteIndex <= 0 {
		return byteIndex
	}
	theBuffer := []byte(src)[:byteIndex]
	// 再转为字符串
	str := string(theBuffer)
	// 转为rune切片
	rs := []rune(str)
	// 返回字符个数
	return len(rs)
}
