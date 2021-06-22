package main

import "fmt"

/*func addOne(x int) {
	x++
}*/

func addOne(x *int) {
	*x ++
}

func main() {
	var n = 7
	fmt.Printf("调用addOne函数前，n的值：%d\n", n)
	// 调用addOne函数
	//addOne(n)
	addOne(&n)
	fmt.Printf("调用addOne函数后，n的值：%d\n", n)
}