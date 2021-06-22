package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 超时限制（秒数）
	const maxSecond = 5
	var (
		// 标志口算完毕
		chFinish = make(chan bool)
		// 标志已超时
		chTimeout = make(chan bool)
	)

	// 生成口算题目
	go func() {
		rand.Seed(time.Now().UnixNano())
		// 产生100以内的随机整数
		a := rand.Intn(100)
		b := rand.Intn(100)
		var input int // 用户输入的计算结果
		r := a + b    // 正确的计算结果
		fmt.Printf("题目：%d + %d = ?\n", a, b)
		fmt.Print("请输入结果：")
		fmt.Scanln(&input)
		// 生成结果
		var cr bool = input == r
		chFinish <- cr
	}()

	// 处理操作超时
	go func() {
		time.Sleep(time.Second * maxSecond)
		chTimeout <- true
	}()

	select {
	case res := <-chFinish:
		if res {
			fmt.Print("\n恭喜你，答对了\n")
		} else {
			fmt.Print("\n噢，答错了\n")
		}
	case <-chTimeout:
		fmt.Print("\n很遗憾，时间到了\n")
	}
	// 关闭通道
	close(chFinish)
	close(chTimeout)
}
