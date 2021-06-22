package main

import (
	"fmt"
	"sync"
	"time"
)

var Total int

var locker = new(sync.Mutex)

func throw() {
	for {
		// 此处开始上锁
		locker.Lock()
		if Total < 1 {
			// 无球可抛时退出
			break
		}
		// 等待一下，抛球需要一定的时间
		time.Sleep(time.Millisecond * 300)
		Total--
		fmt.Printf("剩余%d个球\n", Total)
		// 完成后要解锁
		locker.Unlock()
	}
}

func main() {
	Total = 20
	for i := 0; i < 4; i++ {
		go throw()
	}
	// 暂停一下，等待其他协程完成
	time.Sleep(time.Second * 8)
}
