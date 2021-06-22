package main

import "fmt"

func getThreeInt() (int, int, int) {
	return 100, 300, 500
}

func getSomeStrings() (a, b string) {
	// 为命名的返回值赋值
	a = "Part 1"
	b = "Part 2"
	// 必须用return语句让函数返回
	return
}

func main() {
	var a,b,c = getThreeInt()	
	// 丢弃后两个返回值
	//var a,_,_ = getThreeInt()
	fmt.Printf("a: %d\nb: %d\nc: %d\n", a,b,c)
	//fmt.Printf("a: %d\n", a)

	var s1,s2 = getSomeStrings()
	fmt.Printf("s1: %s\ns2: %s\n", s1,s2)
}