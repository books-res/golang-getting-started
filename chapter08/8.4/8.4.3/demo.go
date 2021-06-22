package main

import "fmt"

func example1() {
	var str string = "春江水暖鸭先知"
	for i,x := range str {
		fmt.Printf("%2d --> %c\n", i,x)
	}
}

func example2() {
	// 数组包含5个元素
	var arr = [5]float32{
		1.00085,
		7.001,
		0.213,
		0.0095,
		205.33,
	}

	for index, element := range arr {
		fmt.Printf("[%d]: %f\n", index, element)
	}	
	// 也可以这样写
	// for index := range arr {
	// 	fmt.Printf("[%d]: %f\n", index, arr[index])
	// }	

	// 丢弃索引
	// for _, element := range arr {
		
	// }
}

func example3() {
	var m = map[rune]string{ 'a': "at", 'b': "bee", 'c': "cut" }
	for key, value := range m {
		fmt.Printf("[%c]: %s\n", key, value)
	}
}

func example4() {
	// 创建通道对象实例
	var ch = make(chan int)

	// 启动新的协程
	go func() {
		// 当代码退出该范围时关闭通道对象
		defer close(ch)
		// 向通道发送内容
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		ch <- 6
	}()

	// 从通道对象中读出所有值
	for v := range ch {
		fmt.Printf("从通道对象中读出：%d\n", v)
	}
}