package main

import (
	"fmt"
)

func main() {
	var sd = `------- 这是标题 -------
一、方案概要
	1、需求分析
	2、技术可行性分析
	3、任务周期分析
二、实施阶段
	a、定义数据
	b、梳理可用资源
	……
						—— 2020年1月5日
----------- 内容底部 ------------
	`
	// 输出到屏幕
	fmt.Println(sd)

	var se string = `Call the "Function"`
	fmt.Println(se)
}