package main

import "fmt"

func main() {
	// 第一种方法：调用make函数
	var m1 = make(map[uint16]string)
	m1[20] = "fly"
	m1[60] = "play"
	fmt.Println(m1)

	// 第二种方法：直接用构造表达式实例化
	var m2 = map[rune]float64{ }
	m2['a'] = 1.0000752
	m2['b'] = -0.00016
	fmt.Println(m2)

	// 在实例化时初始化元素
	var m3 = map[string]uint64{
		"item 1" : 8150,
		"item 2" : 17990,
		"item 3" : 28005,
		"item 4" : 540,
	}
	// 继续添加元素
	m3["item 5"] = 1294
	m3["item 6"] = 290
	m3["item 7"] = 61625
	fmt.Println(m3)
}