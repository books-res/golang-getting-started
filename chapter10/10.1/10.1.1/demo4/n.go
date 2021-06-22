package main

import "fmt"

// 定义接口类型
type music interface {
	play()
	pause()
}

// 实现接口
type popMusic struct { }
func (x popMusic) play() {
	fmt.Println("开始播放流行音乐")
}
func (x popMusic) pause() {
	fmt.Println("暂停播放流行音乐")
}

type classicMusic struct { }
func (x classicMusic) play() {
	fmt.Println("开始播放古典音乐")
}
func (x classicMusic) pause() {
	fmt.Println("暂停播放古典音乐")
}


func main() {
	// 初始化数组实例
	var arr = [2]music{ popMusic{}, classicMusic{} }
	// 调用数组实例中各个元素的方法
	arr[0].play()
	arr[0].pause()
	arr[1].play()
	arr[1].pause()
}