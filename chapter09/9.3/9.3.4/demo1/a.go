package main

import "fmt"

type task interface {
	start()
}

// 下面三个结构体都实现了task接口
type coreTask struct { }

func (t coreTask) start() {
	fmt.Println("启动coreTask任务")
}


type lazyTask struct { }

func (t lazyTask) start() {
	fmt.Println("启动lazyTask任务")
}


type extTask struct { }

func (t extTask) start() {
	fmt.Println("启动extTask任务")
}

/*-----------------------------------------*/
func runTask(t task) {
	// 不管传入的是什么类型的对象
	// 只要对象存在start方法即可
	t.start()
}

/*-----------------------------------------*/

func main() {
	var t1, t2, t3 = coreTask{}, lazyTask{}, extTask{}
	// 调用三次runTask函数
	// 每次传递给函数参数的对象类型不同
	runTask(t1)
	runTask(t2)
	runTask(t3)
}