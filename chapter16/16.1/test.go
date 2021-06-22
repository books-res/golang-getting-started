package main

import "fmt"

func main() {
	// 初始化[]byte实例，长度为0
	var buffer = make([]byte, 0)
	// 向缓冲区添加数据
	buffer = append(buffer, []byte("一二三四五")...)
	buffer = append(buffer, []byte("六七八九十")...)

	// 输出
	fmt.Printf("缓冲区中的数据：%v\n", buffer)
	fmt.Printf("转换为字符串后：%s\n", string(buffer))
}
