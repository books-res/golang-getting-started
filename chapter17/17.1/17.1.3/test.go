package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件
	var file, err = os.Create("myfile")
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	// 写入内容
	var data = []byte{202, 17, 49, 173, 142, 58, 99, 81, 132, 45, 18, 61, 72, 60, 7, 53, 103, 235, 252, 40, 26, 80, 128}
	file.Write(data)
	// 关闭文件
	file.Close()

	// 获取文件信息
	info, err := os.Stat("myfile")
	// 也可以用这种方式获取FileInfo实例
	//info, err := file.Stat()
	if err != nil {
		fmt.Println("错误：", err)
		return
	}
	fmt.Printf("文件名：%s\n", info.Name())
	fmt.Printf("文件大小：%d\n", info.Size())
	// 提取修改时间参数
	var tm = info.ModTime()
	// 依次获取Time对象的年、月、日、时、分、秒部分
	var y, m, d, h, M, s = tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second()
	fmt.Printf("修改时间：%d-%d-%d %d:%d:%d\n", y, m, d, h, M, s)
}
