package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("先等待3秒……")
	time.Sleep(3 * time.Second)
	fmt.Println("噢，再等2秒……")
	time.Sleep(2 * time.Second)
	fmt.Println("等待完毕")
}