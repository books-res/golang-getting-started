package main

import (
	"fmt"
)

func main() {
	if x, y := 3, 5; x + y <= 10 {
		fmt.Println("x、y的和不超过10")
	}
	// 出错，无法访问到变量x
	//fmt.Print(x)
}