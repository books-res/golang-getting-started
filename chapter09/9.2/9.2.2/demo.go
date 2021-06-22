package main

import (
	"fmt"
	"time"
)

type fileInfo struct {
	name string				// 文件名
	size uint64				// 文件大小
	isSysFile bool			// 是否为系统文件
	createTime int64		// 创建时间
}

func main() {
	// 为字段分配默认值
	var x fileInfo
	// 或者
	//var x = fileInfo{ }
	// 或者
	//x := fileInfo{ }
	// 输出各字段的值
	fmt.Printf("文件名：%+v\n", x.name)
	fmt.Printf("文件大小：%d\n", x.size)
	fmt.Printf("是否为系统文件：%t\n", x.isSysFile)
	fmt.Printf("创建时间：%s\n", time.Unix(x.createTime, 0))
	fmt.Println("----------------------------------------------")


	// 使用指针类型必须引用有效的对象实例
	//var px *fileInfo						// nil
	//fmt.Printf("文件名：%s\n", px.name)		// 错误

	// 为字段赋值
	// 方法一
	var y fileInfo
	y.name = "dmd.txt"
	y.isSysFile = false
	y.size = 6955263
	y.createTime = time.Date(2020, 3, 7, 15, 48, 16, 0, time.Local).Unix()
	fmt.Printf("%+v\n", y)

	// 方法二
	var g = fileInfo{
			name: "abc.dat",
			size: 128880, 
			isSysFile: true, 
			createTime: time.Date(2019, 10, 20, 14, 36, 21, 0,  time.Local).Unix(),
		}
	fmt.Printf("%+v\n", g)

	// 忽略部分字段的值
	k := fileInfo {
		name: "dxy.ts",
		size: 3006265,
		createTime: time.Date(2020, 1, 1, 23, 15, 4, 0, time.Local).Unix(),
	}
	fmt.Printf("%+v\n", k)

	// 方法三
	// 省略字段名
	var z = fileInfo{ "test.dat", 1172363, true, time.Now().Unix() }
	fmt.Printf("%+v\n", z)


	// 指针类型
	//var c = fileInfo{ "sys.dll", 23312, true, time.Now().Unix() }
	//var pc *fileInfo = &c
	var pc *fileInfo = &fileInfo{ "sys.dll", 23312, true, time.Now().Unix() }
	fmt.Printf("pc变量所指向的地址：%p\n", pc)
}