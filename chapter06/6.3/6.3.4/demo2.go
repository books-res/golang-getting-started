package main

import (
	"fmt"
	"time"
)

func main() {
	var now = time.Now()
	// 年、月、日
	year, month, day := now.Date()
	// 时、分、秒
	hour, minute, second := now.Clock()

	fmt.Printf("%d年%d月%d日 %d时%d分%d秒", year, month, day, hour, minute, second)
}