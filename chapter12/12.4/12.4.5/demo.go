package main

import (
	"fmt"
	"reflect"
)

// 此函数用于动态构建函数实例（函数体）
// iFn参数为指向函数变量的指针
// 为了简化演示，参数和返回值只支持以下四种类型：
// int32、int64、float32、float64
func initFunc(iFn interface{}) {
	var fnVal = reflect.ValueOf(iFn)
	if fnVal.Kind() != reflect.Ptr {
		fmt.Println("请传递函数变量的内存地址")
		return
	}
	// 取得指针所指向的对象
	fnVal = fnVal.Elem()
	// 函数体逻辑
	var fnBody = func(ip []reflect.Value) (outs []reflect.Value) {
		outs = []reflect.Value{}
		if len(ip) != 2 {
			fmt.Println("函数应该有两个参数")
			return
		}
		// 验证两个参数的类型是否相同
		if ip[0].Kind() != ip[1].Kind() {
			fmt.Println("两个参数的类型不一致")
			return
		}
		// 取出两参数的值
		// 然后进行“+”运算
		var result interface{}
		switch ip[0].Kind() {
		case reflect.Int32:
			a := int32(ip[0].Int())
			b := int32(ip[1].Int())
			result = a + b
		case reflect.Int64:
			a := int64(ip[0].Int())
			b := int64(ip[1].Int())
			result = a + b
		case reflect.Float32:
			a := float32(ip[0].Float())
			b := float32(ip[1].Float())
			result = a + b
		case reflect.Float64:
			a := float64(ip[0].Float())
			b := float64(ip[1].Float())
			result = a + b
		default:
			result = nil
		}
		// 处理返回值
		resval := reflect.ValueOf(result)
		outs = append(outs, resval)
		return
	}

	// 构建函数实例
	funcInst := reflect.MakeFunc(fnVal.Type(), fnBody)
	// 让传递进来的函数变量引用函数实例
	fnVal.Set(funcInst)
}

func main() {
	// 定义不同类型的函数变量
	var (
		addInt32   func(int32, int32) int32
		addInt64   func(int64, int64) int64
		addFloat32 func(float32, float32) float32
		addFloat64 func(float64, float64) float64
	)

	// 初始化以上函数变量
	initFunc(&addInt32)
	initFunc(&addInt64)
	initFunc(&addFloat32)
	initFunc(&addFloat64)

	// 依次调用四个函数
	var a1, a2 int32 = 150, 25
	r1 := addInt32(a1, a2)
	fmt.Printf("%d + %d = %d\n", a1, a2, r1)

	var b1, b2 int64 = 98900000, 45231002
	r2 := addInt64(b1, b2)
	fmt.Printf("%d + %d = %d\n", b1, b2, r2)

	var c1, c2 float32 = 0.0021, 1.0099
	r3 := addFloat32(c1, c2)
	fmt.Printf("%f + %f = %f\n", c1, c2, r3)

	var d1, d2 float64 = -770.00000542, 230.90005
	r4 := addFloat64(d1, d2)
	fmt.Printf("%f + %f = %f\n", d1, d2, r4)
}
