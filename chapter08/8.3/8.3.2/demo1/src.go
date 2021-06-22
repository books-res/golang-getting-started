package main

import "fmt"

func main() {
	var t interface{} = 0.0000012
	switch t.(type) {
	case int:
		fmt.Println("变量t是int类型")
	case float32:
		fmt.Println("变量t是float32类型")
	case float64:
		fmt.Println("变量t是float64类型")
	case rune:
		fmt.Println("变量t是rune类型")
	default:
		fmt.Println("变量t是未知类型")
	}
}