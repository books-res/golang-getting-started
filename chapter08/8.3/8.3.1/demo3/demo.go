package main

import "fmt"

func main() {
	var n = 7
	switch n {
	case 1,3,5,7,9:
		fmt.Println("此整数值是奇数")
	case 0,2,4,6,8:
		fmt.Println("此整数值是偶数")
	}
}