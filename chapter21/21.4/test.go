package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// 增加计数器
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(n int) {
			// 执行完成时将计数器减1
			defer wg.Done()
			fmt.Printf("开始执行第%d个协程\n", n)
			time.Sleep(time.Second * 2)
			fmt.Printf("第%d个协程执行完毕\n", n)
		}(i)
	}

	// 等待上述各协程执行完成
	wg.Wait()
	fmt.Println("所有协程已完成")
}
