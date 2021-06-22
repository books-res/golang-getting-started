package main

import "fmt"

func main() {
	for i := 0; i < 12; i += 2 {
		fmt.Print(i, "  ")
	}

	fmt.Println()
	// 省略初始化子句
	var cc = 'a'

	for ; cc <= 122; cc++ {
		fmt.Printf("%c ", cc)
	}

	fmt.Println()
	// 省略最后一个子句
	for x := 'Z'; x >= 65; {
		fmt.Printf("%c ", x)
		x --
	}

	// 下面语句有错误
	// for i := 1; i < 15 {

	// }

	// var k = 10
	// for k > 0; k-- {

	// }

	// 仅有条件子句
	var n = 5

	for n > 0 {

	}

	// 死循环
	// for {
	// 	fmt.Println("Hello")
	// }
}