package main

import (
	"fmt"
	"strings"
)

func main() {
	var x = 5
	if x > 3 {
		fmt.Print("大于3\n")		// 执行
	}

	if x % 3 == 0 {
		fmt.Println("能被3整除")	// 不执行
	}

	var k string = "check"
	if strings.Contains(k, "ch") {
		fmt.Printf("字符串%s中包含ch\n", k)
	} else {
		fmt.Printf("字符串%s中不包含ch\n", k)
	}

	// 直接用布尔常量值来充当条件表达式
	if true {
		// 此代码会执行
	}

	if false {
		// 此代码不会执行
	}
}