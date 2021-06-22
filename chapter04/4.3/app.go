package main

// 每条import语句只导入一个路径
// import "demo/test1"
// import "demo/test2"
// 合并import语句
import (
	"demo/test1"
	"demo/test2"
)

// 导入时为包分配别名
// import (
// 	ph "demo/test1"
// 	mu "demo/test2"
// )

// 将导入包的成员合并到当前代码中
// import (
// 	. "demo/test1"
// 	. "demo/test2"
// )

func main() {
	// 调用photoplayer包中的函数
	photoplayer.StartPlay()
	photoplayer.StopPlay()
	// 调用musicplayer包中的函数
	musicplayer.Play()
	musicplayer.Pause()

	// 如果导入时为包分配了别名
	// 则访问包成员时应使用别名
	// ph.StartPlay()
	// ph.StopPlay()
	// mu.Play()
	// mu.Pause()

	// 如果import语句中使用.作为别名，访问包成员无需使用包名
	// StartPlay()
	// StopPlay()
	// Play()
	// Pause()
}