package main

import (
	. "./test"
	"fmt"
)

func main() {
	var x = Employee { ID: 11, Name: "小杜", Age: 35, Phone: 13002525123 }

	// 错误：无法访问非公开成员
	//var y = student { ID: 12, Name: "小刘", Age: 21 }
	
	fmt.Printf("%+v\n", x)

	// 错误：字段成员不可访问
	//var z = Printer { current: 0, totalpage: 5, sn: "TD-01" }
}