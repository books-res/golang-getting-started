package main

import "fmt"

func main() {
	fmt.Println("call getFloat：", getFloat())
	fmt.Println("call check(15)：", check(15))
	fmt.Println("call check(18)：", check(18))
}

/*--------- 这两个函数实际上是一样的 ---------*/
func test1() {
	fmt.Println("你好")
	return
}

func test2() {
	fmt.Println("你好")
}


// 有返回值
func getString() string {
	return "Hello"
}

// 第一个return语句的值被返回，其余被忽略
func getFloat() float32 {
	return 0.1
	return 0.2
	return 0.3
}

func check(n int) bool {
	if n % 2 == 0 {
		return true
	}
	return false
}