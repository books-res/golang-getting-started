package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取系统的当前时间
	var ct = time.Now()
	// 获取时间实例各部分的值
	var (
		year = ct.Year()			// 年
		month = ct.Month()			// 月
		day = ct.Day()				// 日
		hour = ct.Hour()			// 时
		minute = ct.Minute()		// 分
		second = ct.Second()		// 秒
	)
	// 向屏幕输出当前时间
	fmt.Printf("当前时间：%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second)
}