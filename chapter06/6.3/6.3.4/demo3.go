package main

import (
	"fmt"
	"time"
)

func main() {
	// 初始化Time对象
	var theTime = time.Date(2015, 12, 3, 0, 0, 0, 0, time.Local)
	// 30小时之后的时间点
	var newTime1 = theTime.Add(30 * time.Hour)
	// 4天之前
	var newTime2 = theTime.Add(-4 * 24 * time.Hour)
	
	// 屏幕输出
	fmt.Printf("原始时间：\n%d-%d-%d %d:%d:%d\n\n",
					theTime.Year(),
					theTime.Month(),
					theTime.Day(),
					theTime.Hour(),
					theTime.Minute(),
					theTime.Second())
	fmt.Printf("30小时后：\n%d-%d-%d %d:%d:%d\n",
					newTime1.Year(),
					newTime1.Month(),
					newTime1.Day(),
					newTime1.Hour(),
					newTime1.Minute(),
					newTime1.Second())	
	fmt.Printf("\n4天以前：\n%d-%d-%d %d:%d:%d",
					newTime2.Year(),
					newTime2.Month(),
					newTime2.Day(),
					newTime2.Hour(),
					newTime2.Minute(),
					newTime2.Second())
}