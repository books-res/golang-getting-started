package main

import "fmt"

func main() {
	var c interface{} = 12
	switch v := c.(type) {
	case string:
		fmt.Printf("字符串：%s\n", v)
	case uint:
		fmt.Printf("无符号整数：%d\n", v)
	case int:
		fmt.Printf("有符号整数：%d\n", v)
	}
}