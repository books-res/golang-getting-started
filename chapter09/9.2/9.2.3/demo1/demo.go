package main

import "fmt"

type test struct {

}

func (o test) doSomething() string {
	return "do nothing"
}
func (o *test) doSomething2() string {
	return "do nothing - 2"
}

func main() {
	var n test
	s := n.doSomething()
	fmt.Println("doSomething方法调用结果：", s)
	s2 := n.doSomething2()
	fmt.Println("通过指针传递：", s2)
}