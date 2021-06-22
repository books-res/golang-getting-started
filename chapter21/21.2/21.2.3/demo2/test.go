package main

func main() {
	// 可行
	var ch1 chan float32 = make(chan float32, 0)
	var ch2 <-chan float32 = ch1
	var ch3 chan<- float32 = ch1

	// 不可行
	var ch4 <-chan int
	var ch5 chan int = ch4
}
