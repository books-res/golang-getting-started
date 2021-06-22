package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 初始化Buffer对象
	var bf bytes.Buffer
	// 写入第一批数据
	var data = []byte{1, 2, 3, 4, 5, 6}
	bf.Write(data)
	// 写入第二批数据
	data = []byte{7, 8, 9, 10}
	bf.Write(data)
	// 写入第三批数据
	data = []byte{11, 12, 13, 14, 15}
	bf.Write(data)

	// 获取全部数据
	var all = bf.Bytes()
	fmt.Printf("缓冲区中的数据：\n%v", all)
}
