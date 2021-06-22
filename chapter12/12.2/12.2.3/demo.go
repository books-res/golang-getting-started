package main

import (
	"fmt"
	"reflect"
)

type test struct {
	m int
	C string
}


func main() {
	// 定义变量
	var vx test
	// 获取类型信息
	tp := reflect.TypeOf(vx)
	// 获取字段成员的数量
	fdnums := tp.NumField()
	// 开始枚举字段
	for x := 0; x < fdnums; x++ {
		fdmember := tp.Field(x)
		// 向屏幕打印文本
		fmt.Printf("字段名称：%s\n", fdmember.Name)
		fmt.Printf("类型：%v\n", fdmember.Type)
		fmt.Printf("程序包路径：%s\n", fdmember.PkgPath)
		fmt.Println()
	}
}