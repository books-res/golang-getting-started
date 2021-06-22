package main

import "fmt"

func main() {
	a := 10
	//for a > 0 {         // 这样无法正常执行
	for ; a > 0; a-- {
		if a == 6 || a == 5 || a == 4 {
			continue
		}
		fmt.Println(a)
		//a--
	}
}