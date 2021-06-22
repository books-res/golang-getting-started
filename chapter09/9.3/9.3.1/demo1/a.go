package demo

type task interface {
	start()						// 无参数，无返回值
	stop() uint16				// 无参数，有返回值
	timeout(long int64) bool	// 有参数，有返回值
}

type runner interface {
	getContext(key string) (uint64, bool)
	// 错误：方法名称重复
	getContext(key int) (uint8, bool)
}

type musicHub interface {
	play(track uint) (stat int)
	_(title string) int		// 方法名称无效
}