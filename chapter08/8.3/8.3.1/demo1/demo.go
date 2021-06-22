package main

import (
	"fmt"
)

func main(){
	var mode = 1
	switch mode {
	case 0:
		fmt.Println("关机状态")
	case 1:
		fmt.Println("开机状态")
	case 2:
		fmt.Println("待机状态")
	}
}