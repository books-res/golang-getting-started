package main

// 结构体
type car struct {
	id uint
	color uint32
}

// 接口
type sender interface {
	writeTo(d string, len int, msg string)
}

// 函数
type otherFunc func(x float32) float32