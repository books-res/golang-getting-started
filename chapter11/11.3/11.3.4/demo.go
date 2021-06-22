package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var r = ring.New(3)
	
	n := r.Len()
	p := r
	c := 'A'
	for i := 0; i < n; i++ {
		p.Value = c
		p = p.Next()
		c++
	}

	var s = ring.New(2)
	
	n = s.Len()
	p = s
	for i := 0; i < n; i++ {
		p.Value = c
		p = p.Next()
		c++
	}

	fmt.Print("链表r：")
	r.Do(func (v interface{}) {
		fmt.Printf("%c ", v)
	})
	fmt.Print("\n链表s：")
	s.Do(func (v interface{}){
		fmt.Printf("%c ", v)
	})

	// 进行链接
	//nr := r.Link(s)
	//若要形成 A-B-C-D-E 的顺序
	// 先把r的指针移到元素C上
	r = r.Move(2)
	nr := r.Link(s)

	fmt.Print("\n链接后：")
	nr.Do(func (v interface{}){
		fmt.Printf("%c ", v)
	})
	fmt.Print("\n")
}