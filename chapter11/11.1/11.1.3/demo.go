package main

import "fmt"

func main() {
	var m = map[string]int{ }

	m["c1"] = 7728
	m["c4"] = 7729

	// 访问key为c3的元素
	// 此key不存在
	//xv := m["c3"]
	//fmt.Println(xv)

	xv, ok := m["c3"]
	if ok {
		fmt.Printf("[\"c3\"] : %d\n", xv)
	} else {
		fmt.Println("\"c3\" 不存在")
	}
}