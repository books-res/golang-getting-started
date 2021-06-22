package main

import (
	"fmt"
	"time"
)

func main() {
	// 2007年6月14日
	var date = time.Date(2007, time.June, 14, 0, 0, 0, 0, time.Local)
	// 获取周工作日
	var weekday = date.Weekday()

	// 用中文表示周工作日名称
	var wdStr string
	switch weekday {
	case time.Sunday:
		wdStr = "星期天"
	case time.Monday:
		wdStr = "星期一"
	case time.Tuesday:
		wdStr = "星期二"
	case time.Wednesday:
		wdStr = "星期三"
	case time.Thursday:
		wdStr = "星期四"
	case time.Friday:
		wdStr = "星期五"
	case time.Saturday:
		wdStr = "星期六"
	default:
		wdStr = "未知"
	}

	// 输出结果
	fmt.Printf("2007-6-14 是 %s", wdStr)
}