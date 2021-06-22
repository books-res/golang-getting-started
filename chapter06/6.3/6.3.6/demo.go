package main

// 导入要使用的包
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 创建新的Timer实例
	var timer = time.NewTimer(5 * time.Second)
	// 创建一个通道实例，用于标识任务已经完成
	var completed = make(chan bool)
	// 退出main函数时关闭通道
	defer close(completed)

	// 在新的协程上执行任务
	go func()  {
		// 随机生成执行任务所需的时间
		// 作用是模拟任务所消耗的时间
		rand.Seed(time.Now().Unix())
		var thelong = rand.Intn(10)
		// 暂停当前协程
		time.Sleep(time.Duration(thelong) * time.Second)
		// 发送信号，表示任务完成
		completed <- true
	}()

	// 判断任务是顺利完成了，还是超时了
	select {
	case <- completed:
		fmt.Println("任务完成")
	case <- timer.C:
		fmt.Println("任务超时")
	}
}