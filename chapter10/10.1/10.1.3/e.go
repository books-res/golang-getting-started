package main

import "fmt"

func main()  {
	var a,b,c = 50,60,70
	var ax = [3]*int{&a, &b, &c}
	// 每个元素都是*int类型
	fmt.Println(ax)

	fmt.Println()

	var d = [3]float32{0.001,0.002,0.003}
	var pd = &d		// 获取数组实例的内存地址
	fmt.Printf("%p\n", pd)
}