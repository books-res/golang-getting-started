package main

import (
	"fmt"
	"time"
)

func main() {
	// 24小时
	var dr = 24 * time.Hour

	// 输出
	fmt.Println("24个小时：")
	fmt.Printf("共有%d分钟\n", int64(dr.Minutes()))
	fmt.Printf("共有%d秒\n", int64(dr.Seconds()))
	fmt.Printf("共有%d毫秒\n", dr.Milliseconds())
	fmt.Printf("共有%d微秒\n", dr.Microseconds())
	fmt.Printf("共有%d纳秒\n", dr.Nanoseconds())
}