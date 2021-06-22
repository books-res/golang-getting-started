package main	//主包
import "fmt"	//导入fmt包

func main() {
	// 8位整数
	var v1 int8
	fmt.Printf("%-15[1]T%[1]v\n", v1)
	
	// 16位整数
	var v2 int16
	fmt.Printf("%-15[1]T%[1]v\n", v2)
	
	// 32位整数
	var v3 int32
	fmt.Printf("%-15[1]T%[1]v\n", v3)
	
	// 64位整数
	var v4 int64
	fmt.Printf("%-15[1]T%[1]v\n", v4)
	
	// 32位浮点数
	var v5 float32
	fmt.Printf("%-15[1]T%[1]v\n", v5)
	
	// 64位浮点数
	var v6 float64
	fmt.Printf("%-15[1]T%[1]v\n", v6)
	
	// 字符串
	var v7 string
	fmt.Printf("%-15[1]T%[1]v\n", v7)
	
	// 单个字符
	var v8 rune
	fmt.Printf("%-15[1]T%[1]v\n", v8)
	
	// 结构体
	type test struct {
		m float32
		n int16
	}
	var v9 test
	fmt.Printf("%-12s%+v\n", "结构体", v9)

	// 接口
	type iTest interface {
		someMethod()
	}
	var v10 iTest
	fmt.Printf("%-13s%v\n", "接口", v10)

	// 指针类型
	var v11 *int
	fmt.Printf("%-13s%v\n", "指针", v11)
}