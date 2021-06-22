package main

import "fmt"

func main() {
	r := 3 + 5 * 2
	fmt.Println(r)
	// 括号改变优先级
	r = (3 + 5) * 2
	fmt.Println(r)

	// 注意：&& 的优先级高于 ||
	b := 3 > 5 || 8 >= 3 && 6 == 0 || 10 <= 0
	fmt.Println(b)

	// 运算符 ! 会先计算
	b = !false || 4 % 2 == 0
	fmt.Println(b)

	b = !(1 < 0 || 4 > 3)
	fmt.Println(b)
}