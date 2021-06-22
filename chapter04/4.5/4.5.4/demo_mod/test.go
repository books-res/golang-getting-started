package hellolib

// 导入标准库中的包
import "fmt"

func SayHello(who string) {
	fmt.Printf("你好，%s\n", who)
}

func SayMorning(who string) {
	fmt.Printf("早上好，%s\n", who)
}