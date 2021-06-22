package main

import (
	"fmt"
)

func main() {
	// 嵌套if语句
	var input int
	fmt.Print("请输入一个整数：")
	fmt.Scan(&input)
	if input >= 80 && input < 100 {
		fmt.Printf("%d在80到100之间\n", input)
	} else if input >= 50 && input < 80 {
		fmt.Printf("%d在50到80之间\n", input)
	} else if input >= 20 && input < 50 {
		fmt.Printf("%d在20到50之间\n", input)
	} else if input >= 0 && input < 20 {
		fmt.Printf("%d在0到20之间\n", input)
	} else {
		fmt.Print("无效范围\n")
	}
}