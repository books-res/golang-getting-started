package main

import (
	"fmt"
	"time"
)

func main() {
	// 2019年7月13日 22:16:08
	var thedate = time.Date(2019, time.July, 13, 22, 16, 8, 0, time.Local)
	fmt.Print(thedate)
	
	// 将Month类型常量转换为int数值
	var n1 = time.August
	var n2 = time.May
	var n3 = time.January
	var n4 = time.November
	fmt.Printf("\nn1: %d\nn2: %d\nn3: %d\nn4: %d", n1, n2, n3, n4)
}