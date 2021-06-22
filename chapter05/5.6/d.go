package main

import (
	"fmt"
)

var x = "EFG"

func main() {
	// 此处覆盖了外部的变量x
	var x string = "XYZ"
	fmt.Print(x)
}