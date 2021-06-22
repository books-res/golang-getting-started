package main

import "fmt"

// func work(fn func(a, b float32) (x, y float32))
// 可以简化为：func work(fn func(float32, float32) (float32, float32))
func work(fn func(float32, float32) (float32, float32)){
	fmt.Println("\nwork函数被调用：")
	if fn != nil {
		r,k := fn(1.0, 0.2)
		fmt.Printf("f1: %f\tf2: %f\n", r,k)
	}
}