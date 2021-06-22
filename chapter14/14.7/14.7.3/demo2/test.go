package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 获取时间戳
	timestamp := time.Now().Unix()
	// 设置随机数种子
	rand.Seed(timestamp)

	// 产生随机数
	for x := 0; x < 5; x++ {
		num := rand.Int()
		fmt.Println(num)
	}
}
