package main

import (
	"fmt"
)

func main() {
	// 初始化映射对象
	myMap := map[string]int {
		"task-01" : 1000,
		"task-02" : 1001,
		"task-03" : 1002,
		"task-04" : 1003,
		"task-05" : 1004,
	}

	for key,val := range myMap {
		fmt.Printf("key: %s\tvalue: %d\n", key, val)
	}
}